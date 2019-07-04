package tokenstore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type TempFileTokenStoreConfig struct {
	GenerateFileName func() string
}

// Config with filename generator based on host name.
// Generator calls os.Exit(1) in case of any errors
var DefaultTempFileTokenStoreConfig = &TempFileTokenStoreConfig{
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

type TempFileTokenStore struct {
	generateFileName func() string
}

// NewTempFileTokenStore creates new instance of TempFileTokenStore
func NewTempFileTokenStore(config *TempFileTokenStoreConfig) *TempFileTokenStore {
	return &TempFileTokenStore{
		generateFileName: config.GenerateFileName,
	}
}

func (store *TempFileTokenStore) Save(token string) error {
	tmpFile, err := ioutil.TempFile(os.TempDir(), store.generateFileName()+"*.auth")
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
