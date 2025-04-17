package command

import (
	"fmt"
	"strings"
)

type Command interface {
	Execute(args []string)
}

var commands = map[string]Command{}

func Register(name string, cmd Command) {
	commands[name] = cmd
}

func HandleInput(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	cmdName := parts[0]
	args := parts[1:]

	if cmd, ok := commands[cmdName]; ok {
		cmd.Execute(args)
	} else {
		fmt.Println("Unknown command:", cmdName)
	}
}

func PrintHelp(args []string) {
	fmt.Println("Available commands:")
	for name := range commands {
		fmt.Println(" -", name)
	}
}

type Func func(args []string)

func (f Func) Execute(args []string) {
	f(args)
}
