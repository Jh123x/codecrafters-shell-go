package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/command"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
	"golang.org/x/term"
)

func main() { // switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Println(err)
		}
	}()

	streamer := parser.NewStreamer(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdStr, err := streamer.GetNextCommand()
		if err != nil {
			fmt.Println("Error reading command")
			continue
		}

		cmd, err := parser.ParseCommands(cmdStr)
		if err != nil {
			fmt.Println("Error parsing command")
			continue
		}

		// Handle Command
		output, err := command.HandleCommand(cmd)
		if err == consts.ErrEXIT {
			break
		}

		fmt.Print("\r\n")

		if len(output) > 0 {
			fmt.Print(output)
		}

		if err != nil {
			if errMsg := err.Error(); len(errMsg) > 0 {
				fmt.Print(errMsg)
			}
		}
		fmt.Print("\r\n")
	}
}
