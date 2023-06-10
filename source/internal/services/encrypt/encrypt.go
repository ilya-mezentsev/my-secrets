package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

type Service struct {
}

func New() Service {
	return Service{}
}

func (s Service) EncryptKey(key string) string {
	return s.hash(key)
}

func (s Service) hash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

func (s Service) EncryptValue(value, password string) (string, error) {
	block, err := aes.NewCipher([]byte(s.hash(password)))
	if err != nil {
		return "", err
	}

	plainText := []byte(value)
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return s.encode(cipherText), nil
}

func (s Service) encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (s Service) DecryptValue(value, password string) (string, error) {
	block, err := aes.NewCipher([]byte(s.hash(password)))
	if err != nil {
		return "", err
	}

	cipherText, err := s.decode(value)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(cipherText))
	iv := ciphertext[:aes.BlockSize]
	cfb := cipher.NewCFBDecrypter(block, iv)

	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func (s Service) decode(value string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}
