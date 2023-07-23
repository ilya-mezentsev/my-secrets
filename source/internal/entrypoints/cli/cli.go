package cli

import (
	"fmt"
	"github.com/howeyc/gopass"
	"my-secrets/internal/services/commands"
	"os"
)

func ProcessCommand(commandsService commands.Service) {
	args := os.Args[1:]
	command, err := commands.Validate(args...)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		showHelp()
	} else {
		processCommand(command, commandsService)
	}
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Printf("\tget key\t- recieve value\n")
	fmt.Printf("\tset key value\t- store value\n")
}

func processCommand(command any, commandsService commands.Service) {
	fmt.Printf("Enter password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		panic("Internal error")
	}

	response := commandsService.Do(command, string(password))

	if response.IsOk {
		fmt.Println(response.Result)
	} else {
		fmt.Printf("Error: %s\n", response.Result)
	}
}
