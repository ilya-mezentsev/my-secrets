package cli

import (
	"fmt"
	"github.com/howeyc/gopass"
	"my-secrets/internal/services/commands"
	"os"
)

func ProcessCommand(commandsService commands.Service) {
	args := os.Args[1:]
	if len(args) < 2 || (args[0] == "set" && len(args) < 3) || !commandsService.IsCommandValid(args[0]) {
		showHelp()
	} else {
		processCommand(args, commandsService)
	}
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Printf("\tget key\t- recieve value\n")
	fmt.Printf("\tset key value\t- store value\n")
}

func processCommand(args []string, commandsService commands.Service) {
	fmt.Printf("Enter password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		panic("Internal error")
	}

	var response commands.Response
	if args[0] == "get" {
		response = commandsService.Get(args[1], string(password))
	} else if args[0] == "set" {
		response = commandsService.Set(args[1], args[2], string(password))
	} else {
		panic("Internal error")
	}

	if response.IsOk {
		fmt.Println(response.Result)
	} else {
		fmt.Printf("Error: %s\n", response.Result)
	}
}
