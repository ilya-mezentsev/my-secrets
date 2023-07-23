package commands

import "errors"

type safeConstructor struct {
	validate func(args ...string) string
	build    func(args ...string) any
}

var commandToValidateFn = map[string]safeConstructor{
	"get": {
		validate: requireLength(1),
		build: func(args ...string) any {
			return GetCommand{key: args[0]}
		},
	},
	"set": {
		validate: requireLength(2),
		build: func(args ...string) any {
			return SetCommand{key: args[0], value: args[1]}
		},
	},
}

func Validate(args ...string) (command any, err error) {
	if len(args) < 1 {
		err = errors.New("no arguments")
		return
	}

	commandName, commandArguments := args[0], args[1:]
	safeConstructor, commandFound := commandToValidateFn[commandName]
	if !commandFound {
		err = errors.New("unexpected command")
		return
	}

	if validationErr := safeConstructor.validate(commandArguments...); validationErr != "" {
		err = errors.New(validationErr)
		return
	}

	command = safeConstructor.build(commandArguments...)

	return
}

func requireLength(expectedLength int) func(args ...string) string {
	return func(args ...string) string {
		if len(args) != expectedLength {
			return "unexpected arguments count"
		} else {
			return ""
		}
	}
}
