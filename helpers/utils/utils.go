package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RootCA struct {
	CertFile      string
	CertPassphase string
}

var (
	once       sync.Once
	RootCAPool *x509.CertPool
)

func LoadRootCA(listRootCA []RootCA) {
	once.Do(func() {
		RootCAPool = x509.NewCertPool()
		for _, item := range listRootCA {
			rootCA, err := ioutil.ReadFile(item.CertFile)
			if err != nil {
				log.Error("reading cert failed : %v", err)
				continue
			}

			RootCAPool.AppendCertsFromPEM(rootCA)
			log.Info("RootCA loaded, file=" + item.CertFile)
		}
	})
}

func StrUnixToTime(s string) time.Time {
	return time.Unix(StrToInt(s), 0)
}

func StrToDuration(strNumber string) time.Duration {
	i := StrToInt(strNumber)
	return time.Duration(i)
}

func IntToStr(inputNum int64) string {
	return strconv.FormatInt(inputNum, 10)
}

func FloatToStr(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

func StrToInt(s string) int64 {

	var number, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return number
}

func StrToFloat(floatStr string) float64 {
	if strings.HasPrefix(floatStr, ".") {
		floatStr = "0" + floatStr
	}
	f, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		return 0.00
	}
	return f
}

func RequestBody(c *gin.Context) string {
	var str string
	buffer, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("RequestBody(), Error:", err.Error())
		return ""
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buffer))
	str = string(buffer)
	return str
}

func SetRequestBody(c *gin.Context, newData string) {
	//set newBuffer
	c.Request.ContentLength = int64(len(newData))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(newData)))
}

func HttpClient(method, url, content_type, data string) (string, error) {
	//Create Instance HTTP Object
	var msg string
	msg = fmt.Sprintf("HttpClient(),***Request***: ,Url:%s ,Method:%s ,Data:%s", url, method, data)
	log.Debug(msg)

	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error("HttpClient(), Error:", err.Error())
		return "", err
	}

	//Set Header Attribute
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", content_type)

	//Create Http Client Instance
	client := &http.Client{Timeout: time.Second * 60}
	res, err := client.Do(req)
	if err != nil {
		log.Error("HttpClient(), Error:", err.Error())
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		err := errors.New(fmt.Sprintf("%d %s", res.StatusCode, http.StatusText(res.StatusCode)))
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("HttpClient(), Error:", err.Error())
		return "", err
	}

	result := string(body)

	msg = fmt.Sprintf("HttpClient(),***Response***: ,Data:%s", result)
	log.Debug(msg)

	return result, nil

}

func HttpsClientOpt(
	method string,
	url string,
	headers map[string]string,
	data string) (string, error) {
	//Create Instance HTTP Object
	var msg string
	msg = fmt.Sprintf("HttpsClient(),***Request*** :Url:%s, Method:%s, Data:%s", url, method, data)
	log.Debug(msg)

	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error("HttpsClient(), Error:", err.Error())
		return "", err
	}

	//Set Header Attribute
	req.Header.Set("Cache-Control", "no-cache")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//Create Http Client Instance
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: RootCAPool, InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}
	res, err := client.Do(req)

	if err != nil {
		log.Error("HttpsClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		errDesc := ""
		if res != nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			errDesc += string(body)
		}
		err := errors.New(errDesc)
		log.Error("HttpsClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("HttpClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	result := string(body)

	msg = fmt.Sprintf("HttpClient(),***Response*** :Data:%s", result)
	log.Info(msg)

	return result, nil

}

func HttpsClientOpt2(
	method string,
	url string,
	headers map[string]string,
	data string) (string, error) {
	//Create Instance HTTP Object
	var msg string
	msg = fmt.Sprintf("HttpsClient(),***Request***: Url:%s, Method:%s ,Data:%s", url, method, data)
	log.Debug(msg)

	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error("HttpsClient(), Error:", err.Error())
		return "", err
	}

	//Set Header Attribute
	req.Header.Set("Cache-Control", "no-cache")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// create a Certificate pool to hold one or more CA certificates
	rootCAPool := x509.NewCertPool()
	RootCertificatePath := "D:/WorkSpace/GoWorkspace/src/ktb-dialer-middleware/rlcs_cer.cer"

	// read minica certificate (which is CA in our case) and add to the Certificate Pool
	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")

	client := &http.Client{Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{RootCAs: rootCAPool},
		}}
	res, err := client.Do(req)

	if err != nil {
		log.Error("HttpsClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		errDesc := ""
		if res != nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			errDesc += string(body)
		}
		err := errors.New(errDesc)
		log.Error("HttpsClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("HttpClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	result := string(body)

	msg = fmt.Sprintf("HttpClient(),***Response***: ,Data:%s", result)
	log.Info(msg)

	return result, nil

}

func HttpClientOpt(
	method string,
	url string,
	headers map[string]string,
	data string) (string, error) {
	//Create Instance HTTP Object
	var msg string
	msg = fmt.Sprintf("HttpClient(),***Request***: ,Url:%s ,Method:%s ,Data:%s", url, method, data)
	log.Debug(msg)

	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Error("HttpClient(), Error:", err.Error())
		return "", err
	}

	//Set Header Attribute
	req.Header.Set("Cache-Control", "no-cache")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//Create Http Client Instance
	client := &http.Client{Timeout: time.Second * 60}
	res, err := client.Do(req)

	if err != nil {
		log.Error("HttpClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		errDesc := ""
		if res != nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			errDesc += string(body)
		}
		err := errors.New(errDesc)
		log.Error("HttpClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error("HttpClient(), ***Response***, Error:", err.Error())
		return "", err
	}

	result := string(body)

	msg = fmt.Sprintf("HttpClient(),***Response***: ,Data:%s", result)
	log.Info(msg)

	return result, nil

}

func StructToMapInterface(dataStruct interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(dataStruct)
	json.Unmarshal(inrec, &inInterface)

	return inInterface
}

func StructToJsonString(data interface{}) string {
	var result string
	buffer, err := json.Marshal(data)
	if err != nil {
		result = ""
	}
	result = string(buffer)
	return result
}

func JsonStringToStruct(dataJson string, data interface{}) (err error) {
	err = json.Unmarshal([]byte(dataJson), &data)
	return err
}

func StrToDate(strDate string) time.Time {
	strDate = strings.ReplaceAll(strDate, "T", " ")
	strDate = strings.ReplaceAll(strDate, "Z", "")

	pDate, err := time.Parse("2006-01-02 15:04:05", strDate)

	if err != nil {
		return time.Now()
	}
	return pDate
}

func DateDiff(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}

func TrimLF(str string) string {
	str = strings.ReplaceAll(str, "\r\n", " ")
	return str
}
