package ebook

import (
	"os"
	"strings"
	"testing"
)

func TestSaveImageSpecifyDirectory(t *testing.T) {
	//url := "https://saitodev.co/site/image/soycmsbanner.jpg"
}

func TestSaveImage(t *testing.T) {
	url := "https://saitodev.co/site/image/soycmsbanner.jpg"
	err := SaveImage(url)
	if err != nil {
		t.Error("画像の保存で失敗しました")
	}

	n := strings.LastIndex(url, "/")
	filename := url[n+1:]

	_, err = os.Stat(filename)
	if err != nil {
		t.Error("ファイルが存在していませんでした")
	}

	err = os.Remove(filename)
	if err != nil {
		t.Error("ファイルの削除で失敗しました")
	}
}
