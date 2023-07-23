package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"my-secrets/internal/repositories/secret"
	"my-secrets/internal/services/encrypt"
)

type (
	Service struct {
		secretsRepository secret.IRepository
		cipher            encrypt.IService
	}

	Response struct {
		IsOk   bool
		Result string
	}
)

func New(secretsRepository secret.IRepository, cipher encrypt.IService) Service {
	return Service{
		secretsRepository: secretsRepository,
		cipher:            cipher,
	}
}

func (s Service) Do(command any, password string) Response {
	switch command.(type) {
	case GetCommand:
		return s.get(command.(GetCommand), password)

	case SetCommand:
		return s.set(command.(SetCommand), password)

	default:
		panic("internal error")
	}
}

func (s Service) get(command GetCommand, password string) Response {
	key := command.key

	hashedKey := s.cipher.EncryptKey(key)
	keyExists, err := s.secretsRepository.Exists(hashedKey)
	if err != nil {
		log.Errorf("Unable to read value for key: %s; got error: %v", key, err)

		return Response{
			IsOk:   false,
			Result: "Got error while reading value.",
		}
	}

	if !keyExists {
		return Response{
			IsOk:   false,
			Result: fmt.Sprintf("Key '%s' is not found.", key),
		}
	}

	value, err := s.secretsRepository.Fetch(hashedKey)
	if err != nil {
		log.Errorf("Unable to read value for key: %s; got error: %v", key, err)

		return Response{
			IsOk:   false,
			Result: "Got error while reading value.",
		}
	}

	decryptedValued, err := s.cipher.DecryptValue(value, password)
	if err != nil {
		log.Errorf("Unable to decrypt value for key: %s; got error: %v", key, err)

		return Response{
			IsOk:   false,
			Result: "Got error while decrypting value. Please, check password and try again",
		}
	}

	return Response{
		IsOk:   true,
		Result: decryptedValued,
	}
}

func (s Service) set(command SetCommand, password string) Response {
	key, value := command.key, command.value

	hashedKey := s.cipher.EncryptKey(key)
	encryptedValue, err := s.cipher.EncryptValue(value, password)
	if err != nil {
		log.Errorf("Unable to encrypt value, got error: %v", err)

		return Response{
			IsOk:   false,
			Result: "Unable to encrypt passed value",
		}
	}

	err = s.secretsRepository.Save(hashedKey, encryptedValue)
	if err != nil {
		log.Errorf("Unable to save secret value, got error: %v", err)

		return Response{
			IsOk:   false,
			Result: "Unable to save passed value",
		}
	}

	return Response{
		IsOk:   true,
		Result: "Value successfully saved",
	}
}
