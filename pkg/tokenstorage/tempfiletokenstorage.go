package tokenstorage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const fileNamePrefix = "echo-cli.auth."

type TempFileTokenStorageConfig struct {
	GenerateFileName func() (string, error)
}

// Config with filename generator based on host name.
var DefaultTempFileTokenStorageConfig = &TempFileTokenStorageConfig{
	GenerateFileName: generateFileNameByHost,
}

// generateFileNameByHost creates filename based on host name
func generateFileNameByHost() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", errors.Wrap(err, "Failed to get hostname")
	}

	return name, nil
}

type TempFileTokenStorage struct {
	generateFileName func() (string, error)
	tokenStorePath   string
}

// NewTempFileTokenStorage creates new instance of TempFileTokenStorage
func NewTempFileTokenStorage(config *TempFileTokenStorageConfig) *TempFileTokenStorage {
	return &TempFileTokenStorage{
		generateFileName: config.GenerateFileName,
		tokenStorePath:   os.TempDir(),
	}
}

func (storage *TempFileTokenStorage) Save(token string) error {
	err := storage.deleteExistingTokens()
	if err != nil {
		return errors.Wrap(err, "Failed to delete old token")
	}

	filename, err := storage.generateFileName()
	if err != nil {
		return errors.Wrap(err, "Failed generate filename")
	}

	tmpFile, err := ioutil.TempFile(storage.tokenStorePath, fileNamePrefix+filename+"*.auth")
	if err != nil {
		return errors.Wrap(err, "Failed to create temporary file with aut token")
	}

	fmt.Printf("Auth file saved to %v\n", tmpFile.Name())

	_, err = tmpFile.WriteString(token)
	if err != nil {
		return errors.Wrap(err, "Failed to save auth token")
	}

	return nil
}

func (storage *TempFileTokenStorage) deleteExistingTokens() error {
	result, err := ioutil.ReadDir(os.TempDir())
	if err != nil {
		return errors.Wrap(err, "Failed to read temp dir")
	}

	for _, file := range result {
		if file.IsDir() || !strings.HasPrefix(file.Name(), fileNamePrefix) {
			continue
		}
		err := os.Remove(filepath.Join(storage.tokenStorePath, file.Name()))
		if err != nil {
			return errors.Wrap(err, "Failed to delete old auth token")
		}
	}

	return nil
}

func (storage *TempFileTokenStorage) Read() (string, error) {
	result, err := ioutil.ReadDir(storage.tokenStorePath)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read temp dir")
	}

	for _, file := range result {
		if file.IsDir() || !strings.HasPrefix(file.Name(), fileNamePrefix) {
			continue
		}
		tokenAsBuffer, err := ioutil.ReadFile(filepath.Join(storage.tokenStorePath, file.Name()))
		if err != nil {
			return "", errors.Wrap(err, "Failed to read token from file")
		}

		return string(tokenAsBuffer), nil
	}

	return "", errors.WithStack(ErrTokenNotFound)
}
