package helpers

import (
	"crypto/sha256"
	"io/ioutil"
	"path/filepath"
)

func GetAbsPathAndReadFile(path string) ([]byte, error) {

	// Create absolute path to config file
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Read Config File
	fileContent, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		return nil, err
	}

	return fileContent, nil

}

func CalcFileChecksum(data []byte) ([]byte, error) {

	sum := sha256.Sum256(data)

	return sum[:], nil

}
