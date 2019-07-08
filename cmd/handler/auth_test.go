package handler

import (
	"testing"

	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

const ConstToken string = "sec_token"

type TokenStorageMock struct {
	mock.Mock
}

func (d *TokenStorageMock) Save(token string) error {
	args := d.Called(token)
	return args.Error(0)
}

func (d *TokenStorageMock) Read() (string, error) {
	args := d.Called()
	return args.String(0), args.Error(1)
}

func mockErrorHandler(err error) error {
	return err
}

func Test_Handle_OnStorageWithToken_ReturnsToken(t *testing.T) {

	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return(ConstToken, nil)
	auth := NewAuthHandler(storage, mockErrorHandler)

	err := auth.Handle(nil, nil)

	assert.Nil(t, err)
}

func Test_GetToken_OnStorageWithoutToken_ReturnsErr(t *testing.T) {

	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return("", tokenstorage.ErrTokenNotFound)
	auth := NewAuthHandler(storage, mockErrorHandler)

	err := auth.Handle(nil, nil)

	assert.NotNil(t, err)
	assert.Equal(t, tokenstorage.ErrTokenNotFound.Error(), err.Error())
}
