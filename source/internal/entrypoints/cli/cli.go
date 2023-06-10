package cli

import (
	"my-secrets/internal/services/commands"
	"os"
)

func ProcessCommand(commandsService commands.Service) {
	// todo parse and validate args
	_ = os.Args[1:]
}
