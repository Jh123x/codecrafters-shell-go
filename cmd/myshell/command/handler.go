package command

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func HandleCommand(command *parser.Command) (res string, err error) {
	if command == nil {
		return "", consts.ErrEmptyCommand
	}

	stdout, err := handleCommand(command)
	if command.Link != nil {
		return handleLink(command.Link, stdout, err)
	}

	return strings.ReplaceAll(stdout, "\n", "\r\n"), err
}

func handleLink(link *parser.Link, stdout string, stderr error) (string, error) {
	if link == nil || len(link.Args) == 0 {
		return "", consts.ErrUnexpectedLinkValue
	}

	cmd := link.Args
	switch link.Type {
	case parser.LinkTypeStdout:
		if err := WriteToFile(cmd[0], stdout); err != nil {
			fmt.Println("error while writing stdout to file", err)
			return "", err
		}
		if stderr == nil {
			fmt.Println("stderr is empty in stdout redirection flow")
			return "", nil
		}
		return "", stderr

	case parser.LinkTypeStderr:
		errTxt := ""
		if stderr != nil {
			fmt.Println("stderr is empty in stderr redirection flow")
			errTxt = stderr.Error()
		}

		if err := WriteToFile(cmd[0], errTxt); err != nil {
			fmt.Println("error while writing stderr to file", err)
			return "", err
		}

		return stdout, nil
	case parser.LinkTypeAppendStdout:
		if err := AppendToFile(cmd[0], stdout); err != nil {
			fmt.Println("error while appending to file", err)
			return "", err
		}
		if stderr == nil {
			fmt.Println("stderr is empty in stdout redirection append flow")
			return "", nil
		}
		return "", stderr
	case parser.LinkTypeAppendStderr:
		errTxt := ""
		if stderr != nil {
			fmt.Println("stderr is empty in stderr redirection append flow")
			errTxt = stderr.Error()
		}

		if err := AppendToFile(cmd[0], errTxt); err != nil {
			fmt.Println("error while appending to file", err)
			return "", err
		}

		return stdout, nil
	default:
		fmt.Println("invalid link type")
		return "", consts.ErrUnsupportedLinkType
	}
}

func handleCommand(command *parser.Command) (string, error) {
	switch command.Command {
	case consts.ECHO:
		return Echo(command.Args)
	case consts.EXIT:
		return Exit(command.Args)
	case consts.TYPE:
		return Type(command.Args)
	case consts.PWD:
		return Pwd()
	case consts.CD:
		return ChangeDir(command.Args)

	// For debugging command struct.
	case consts.INSPECT:
		return Inspect(command)

	// Executable files.
	default:
		return DefaultCommand(command)
	}
}
