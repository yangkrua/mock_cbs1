package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func MoveFile(src, dst string) error {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		dir := filepath.Dir(dst)
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			log.Error(err.Error())
		}
	}

	err := os.Rename(src, dst)
	if err != nil {
		log.Error("MoveFile(), Error:", err.Error())
	}
	return err
}

func ReadAllFile(src string) string {
	var data string = ""
	file, err := os.Open(src)
	if err != nil {
		log.Error("ReadAllFile(), Error:", err.Error())
		return data
	}

	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error("ReadAllFile(), Error:", err.Error())
		return data
	}
	data = string(buffer)
	return data
}

func WriteFile(src, data string) bool {
	var ok bool = true

	dir := filepath.Dir(src)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			log.Error(err.Error())
		}
	}

	f, err := os.OpenFile(src, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		ok = false
		return ok
	}

	defer f.Close()

	if _, err = f.WriteString(data + "\n"); err != nil {
		ok = false
		return ok
	}

	return ok
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
