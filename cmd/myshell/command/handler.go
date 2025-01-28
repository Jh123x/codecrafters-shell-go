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
	if link == nil || link.LinkedCommand == nil || link.LinkedCommand.Command == "" {
		return "", consts.ErrUnexpectedLinkValue
	}

	cmd := link.LinkedCommand
	switch link.Type {
	case parser.LinkTypeStdout:
		if err := WriteToFile(cmd.Command, stdout); err != nil {
			return "", err
		}
		if stderr == nil {
			return "", nil
		}
		return "", stderr

	case parser.LinkTypeStderr:
		if stderr == nil {
			return stdout, nil
		}

		if err := WriteToFile(cmd.Command, stderr.Error()); err != nil {
			return "", err
		}

		return "", stderr
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
	default:
		return DefaultCommand(command)
	}
}
