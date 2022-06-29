package helpers

import (
	"bytes"
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

func CompareData(source []byte, target []byte) (bool, error) {

	sourceSha256, err := CalcFileChecksum(source)
	if err != nil {
		return false, err
	}

	targetSha256, err := CalcFileChecksum(target)
	if err != nil {
		return false, err
	}

	res := bytes.Compare(sourceSha256, targetSha256)
	if res != 0 {
		return false, nil
	}

	return true, nil
}

func CalcFileChecksum(data []byte) ([]byte, error) {

	sum := sha256.Sum256(data)

	return sum[:], nil

}
