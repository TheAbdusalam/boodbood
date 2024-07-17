package main

import (
	"os"
	"path/filepath"
)

var home string = os.Getenv("HOME")
var actionFile string = filepath.Join(home, ".config/boodbood/actions/run.sh")
var storeFilePath string = filepath.Join(home, ".config/boodbood/store.txt")

func init() {
	if _, err := os.Stat(actionFile); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(actionFile), 0755)
		_ = os.WriteFile(actionFile, []byte(""), 0644)
	}

	if _, err := os.Stat(storeFilePath); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(storeFilePath), 0755)
		_ = os.WriteFile(storeFilePath, []byte(""), 0644)
	}
}

func main() {
	bd := NewBoodBood(actionFile)
	bd.Navigate(os.Args[1:])

}
