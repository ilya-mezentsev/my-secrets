package encrypt

type IService interface {
	EncryptKey(key string) string
	EncryptValue(value, password string) (string, error)
	DecryptValue(value, password string) (string, error)
}
