package secret

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

type Repository struct {
	basePath string
}

func New(basePath string) Repository {
	return Repository{basePath: basePath}
}

func (r Repository) Save(key, value string) error {
	filePath := path.Join(r.basePath, key)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Errorf("Unable to close file: %v; path: %s", err, filePath)
		}
	}()

	_, err = f.Write([]byte((value)))
	return err
}

func (r Repository) Fetch(key string) (string, error) {
	filePath := path.Join(r.basePath, key)
	content, err := os.ReadFile(filePath)

	return string(content), err
}
