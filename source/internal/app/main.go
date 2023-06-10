package app

import (
	"fmt"
	"my-secrets/internal/entrypoints/cli"
	"my-secrets/internal/repositories/secret"
	"my-secrets/internal/services/commands"
	"my-secrets/internal/services/encrypt"
	"os"
)

func Main() {
	// todo read from env
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Unable to read user home dir: %v\n", err)
	} else {
		processCommand(userHomeDir)
	}
}

func processCommand(userHomeDir string) {
	commandsService := commands.New(
		secret.New(userHomeDir),
		encrypt.New(),
	)

	// todo somehow change entrypoints?
	cli.ProcessCommand(commandsService)
}
