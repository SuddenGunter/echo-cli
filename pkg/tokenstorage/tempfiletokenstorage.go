package tokenstorage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type TempFileTokenStorageConfig struct {
	GenerateFileName func() string
}

// Config with filename generator based on host name.
// Generator calls os.Exit(1) in case of any errors
var DefaultTempFileTokenStorageConfig = &TempFileTokenStorageConfig{
	GenerateFileName: generateFileNameByHost,
}

// generateFileNameByHost creates filename based on host name
// calls os.Exit(1) in case of any errors
func generateFileNameByHost() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return name
}

type TempFileTokenStorage struct {
	generateFileName func() string
}

// NewTempFileTokenStorage creates new instance of TempFileTokenStorage
func NewTempFileTokenStorage(config *TempFileTokenStorageConfig) *TempFileTokenStorage {
	return &TempFileTokenStorage{
		generateFileName: config.GenerateFileName,
	}
}

func (storage *TempFileTokenStorage) Save(token string) error {
	tmpFile, err := ioutil.TempFile(os.TempDir(), storage.generateFileName()+"*.auth")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}
	fmt.Printf("Log saved to %v%v", os.TempDir(), tmpFile.Name())

	_, err = tmpFile.WriteString(token)
	if err != nil {
		return err
	}

	return nil
}
