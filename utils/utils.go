package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func LoadFileString(path string) string {
	wd, _ := os.Getwd()
	file, err := ioutil.ReadFile(wd + path)
	if err != nil {
		log.Fatalln(err)
	}
	return string(file)
}
