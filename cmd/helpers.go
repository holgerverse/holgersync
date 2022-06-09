package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func getAbsPathAndReadFile(path string) []byte {

	// Create absolute path to config file
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	// Read Config File
	fileContent, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return fileContent

}
