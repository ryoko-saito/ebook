package ebook

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SaveImageSpecifyDirectory(url string, dir string) error {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	//ワーキングディレクトリを指定のディレクトリに変更
	_, err = os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	err = os.Chdir(currentDir + "/" + dir)
	if err != nil {
		return err
	}

	err = SaveImage(url)
	if err != nil {
		return err
	}

	//ワーキングディレクトリを戻す
	err = os.Chdir(currentDir)

	return err
}

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
