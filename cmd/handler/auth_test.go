package handler

import (
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/require"

	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"

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

const MockErrorHandlerCalled string = "mock error handler called"

func mockErrorHandler(err error) error {
	return errors.WithMessage(err, MockErrorHandlerCalled)
}

func Test_Handle_OnStorageWithToken_ReturnsToken(t *testing.T) {
	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return(ConstToken, nil)
	auth := NewAuthHandler(storage, mockErrorHandler)

	err := auth.Handle(nil, nil)

	require.Nil(t, err)
}

func Test_GetToken_OnStorageWithoutToken_ErrorHandlerCalled(t *testing.T) {
	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return("", tokenstorage.ErrTokenNotFound)
	auth := NewAuthHandler(storage, mockErrorHandler)

	err := auth.Handle(nil, nil)
	startsWith := func() bool {
		return strings.HasPrefix(err.Error(), MockErrorHandlerCalled)
	}

	require.NotNil(t, err)
	require.Condition(t, startsWith)
	require.Equal(t, tokenstorage.ErrTokenNotFound.Error(), errors.Cause(err).Error())
}

func Test_GetToken_WhenTokenExists_ReturnsToken(t *testing.T) {
	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return(ConstToken, nil)
	auth := NewAuthHandler(storage, mockErrorHandler)

	auth.token = ConstToken
	token, err := auth.GetToken()

	require.Nil(t, err)
	require.Equal(t, ConstToken, token)
}

func Test_GetToken_WhenTokenEmpty_ReturnsErr(t *testing.T) {
	storage := new(TokenStorageMock)
	storage.On("Save", ConstToken).Return(nil)
	storage.On("Read").Return(ConstToken, nil)
	auth := NewAuthHandler(storage, mockErrorHandler)

	auth.token = ""
	_, err := auth.GetToken()

	require.NotNil(t, err)
}
