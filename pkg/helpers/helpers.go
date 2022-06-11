package helpers

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetAbsPathAndReadFile(path string) string {

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

func CalcFileChecksum(path string) [32]byte {

	sum := sha256.Sum256([]byte(GetAbsPathAndReadFile(path)))

	return sum

}
