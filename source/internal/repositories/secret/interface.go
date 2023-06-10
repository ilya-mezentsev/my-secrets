package secret

type IRepository interface {
	Save(key, value string) error
	Fetch(key string) (string, error)
}
