package main

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"path/filepath"
)

func getAbsPathAndReadFile(path string) string {

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

	return string(fileContent)

}

func calcFileChecksum(path string) [32]byte {

	sum := sha256.Sum256([]byte(getAbsPathAndReadFile(path)))

	return sum

}
