package ebook

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSaveImageSpecifyDirectory(t *testing.T) {
	url := "https://saitodev.co/site/image/soycmsbanner.jpg"
	dir := "folder"
	err := SaveImageSpecifyDirectory(url, dir)
	if err != nil {
		t.Error("指定の箇所に画像を保存できませんでした")
	}
	_, err = os.Stat(dir)
	if err != nil {
		t.Error("ディレクトリが存在しませんでした")
	}

	n := strings.LastIndex(url, "/")
	filename := url[n+1:]

	//実行しているコードと同階層にファイルがある場合はエラー
	_, err = os.Stat(filename)
	if err == nil {
		t.Error("意図していないディレクトリにファイルが作成されています")
	}

	currentDir, _ := filepath.Abs(".")
	_ = os.Chdir(currentDir + "/" + dir)
	_, err = os.Stat(filename)
	if err != nil {
		t.Error("指定の箇所にファイルが作成されていません")
	}

	err = os.Remove(filename)
	if err != nil {
		t.Error("ファイルの削除で失敗しました")
	}

	_ = os.Chdir(currentDir)
	err = os.Remove(dir)
	if err != nil {
		t.Error("ディレクトリの削除で失敗しました")
	}
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
