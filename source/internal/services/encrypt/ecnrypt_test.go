package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = New()

const (
	password       = "password"
	testValue      = "foo42"
	encryptedValue = "X5HPz3A="
)

func TestService_EncryptKey(t *testing.T) {
	assert.Equal(t, s.EncryptKey(testValue), "097e2944de65e7a83f7527114eb69745")
}

func TestService_EncryptValue_Success(t *testing.T) {
	v, err := s.EncryptValue(testValue, password)

	assert.Nil(t, err)
	assert.Equal(t, encryptedValue, v)
}

func TestService_DecryptValue_Success(t *testing.T) {
	v, err := s.DecryptValue(encryptedValue, password)

	assert.Nil(t, err)
	assert.Equal(t, testValue, v)
}

func TestService_DecryptValue_WrongInput(t *testing.T) {
	v, err := s.DecryptValue("trash", password)

	assert.NotNil(t, err)
	assert.Empty(t, v)
}
