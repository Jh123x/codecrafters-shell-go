package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	commandStr, err := bufio.NewReader(os.Stdin).ReadString('\n')

	// Handle Command
	if err != nil {
		fmt.Println("Error reading command")
	}

	if err := command.HandleCommand(os.Stdout, commandStr, nil); err != nil {
		fmt.Print("error handling command: ", err)
	}
}
