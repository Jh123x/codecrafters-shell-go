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
		fmt.Println(err)
		return
	}
	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Println(err)
		}
	}()

	streamer := parser.NewStreamer(os.Stdin)
	for {
		fmt.Print("\r$ ")
		cmdStr, err := streamer.GetNextCommand()
		if err != nil {
			fmt.Println("Error reading command")
			continue
		}

		cmd, err := parser.ParseCommands(cmdStr)
		if err != nil {
			fmt.Println("Error reading command")
			continue
		}

		// Handle Command
		output, err := command.HandleCommand(cmd)
		if err == consts.ErrEXIT {
			break
		}

		if len(output) > 0 {
			fmt.Printf("%s\r", output)
		}

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
	}
}
