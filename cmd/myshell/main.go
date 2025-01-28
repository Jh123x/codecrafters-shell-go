package main

import (
	"fmt"
	"os"
	"strings"

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
			fmt.Print(fixStrPrinting(output))
		}

		if err != nil {
			if errMsg := err.Error(); len(errMsg) > 0 {
				fmt.Print(fixStrPrinting(errMsg))
			}
		}
	}
}

func fixStrPrinting(val string) string {
	return strings.ReplaceAll(strings.ReplaceAll(val, "\r\n", "\n"), "\n", "\r\n") + "\r\n"
}
