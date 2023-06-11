package app

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"my-secrets/internal/entrypoints/cli"
	"my-secrets/internal/repositories/secret"
	"my-secrets/internal/services/commands"
	"my-secrets/internal/services/encrypt"
	"os"
	"path"
)

var (
	appDirPath  string
	secretsPath string
)

const appName = ".my-secrets"

func init() {
	// todo read from env
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to read user home dir: %v\n", err)
		os.Exit(1)
	}

	appDirPath = path.Join(userHomeDir, appName)
	secretsPath = path.Join(appDirPath, "secrets")

	ensureDirExists(appDirPath)
	ensureDirExists(secretsPath)
	setupLogging()
}

func ensureDirExists(dirPath string) {
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Unable to create application directory: %v\n", err)
			os.Exit(1)
		}
	}
}

func setupLogging() {
	if os.Getenv("VERBOSE_SECRETS") == "Y" {
	} else {
		log.SetOutput(io.Discard)
	}
}

func Main() {
	commandsService := commands.New(
		secret.New(secretsPath),
		encrypt.New(),
	)

	// todo change entrypoints somehow?
	cli.ProcessCommand(commandsService)
}
