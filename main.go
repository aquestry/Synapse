package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aquestry/synapse/command"
)

func main() {
	// Register all commands
	command.Register("list", &command.ListCommand{})
	command.Register("help", command.Func(command.PrintHelp))

	fmt.Println("Docker CLI (type 'help' for a list of commands)")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		input := line[:len(line)-1]
		if input == "exit" || input == "quit" {
			fmt.Println("Exiting.")
			return
		}

		command.HandleInput(input)
	}
}
