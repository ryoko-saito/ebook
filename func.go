package ebook

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func SaveImage(url string) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	n := strings.LastIndex(url, "/")
	p := url[n+1:]

	file, err := os.Create(p)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}

	return nil
}
