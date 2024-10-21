package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
  func ForwardToServer(host string, c *gin.Context) bool
*/
func ForwardToServer(host string, c *gin.Context, data string) (string, error) {
	url, err := url.Parse(host)
	if err != nil {
		log.Error("ForwardToServer(), Error: NewRequest...,", err.Error())
	}

	proxyReq, err := http.NewRequest(c.Request.Method, url.String(), bytes.NewBuffer([]byte(data)))
	for header, values := range c.Request.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	client := &http.Client{Timeout: time.Second * 30}
	proxyRes, err := client.Do(proxyReq)

	if proxyRes.StatusCode != http.StatusOK || err != nil {
		log.Error("ForwardToServer(), Error: NewRequest...,", err.Error())
		return "", err
	}
	defer proxyRes.Body.Close()

	buffer, _ := ioutil.ReadAll(proxyRes.Body)
	result := string(buffer)

	return result, nil
}
