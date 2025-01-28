package command

import (
	"fmt"

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

	return stdout, err
}

func handleLink(link *parser.Link, stdout string, stderr error) (string, error) {
	if link == nil || len(link.Args) == 0 {
		return "", consts.ErrUnexpectedLinkValue
	}

	cmd := link.Args
	switch link.Type {
	case parser.LinkTypeStdout:
		if err := WriteToFile(cmd[0], stdout); err != nil {
			return "", err
		}
		if stderr == nil {
			return "", nil
		}
		return "", stderr

	case parser.LinkTypeStderr:
		errTxt := ""
		if stderr != nil {
			errTxt = stderr.Error()
		}

		if err := WriteToFile(cmd[0], errTxt); err != nil {
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
	case consts.Inspect:
		return Inspect(command)
	default:
		return DefaultCommand(command)
	}
}
