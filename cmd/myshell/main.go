package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Print("$ ")

	// Wait and parse user input
	cmd, args, err := parser.ParseFromReader(os.Stdin)
	if err != nil {
		fmt.Println("Error reading command")
	}

	// Handle Command
	output, err := command.HandleCommand(cmd, args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(output)
}
