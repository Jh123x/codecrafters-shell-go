package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func main() {
	for {
		fmt.Print("$ ")
		cmd, err := parser.ParseFromReader(os.Stdin)
		if err != nil {
			fmt.Println("Error reading command")
		}

		// Handle Command
		output, err := command.HandleCommand(cmd)
		if err == consts.ErrEXIT {
			return
		}
		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		fmt.Print(output)
	}
}
