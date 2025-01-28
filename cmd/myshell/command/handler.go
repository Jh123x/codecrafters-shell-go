package command

import (
	"io/fs"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func HandleCommand(command *parser.Command) (string, error) {
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
	if link == nil {
		return "", consts.ErrNilLink
	}

	linkCommand := link.LinkedCommand
	switch link.Type {
	case parser.LinkTypeStdout:
		if linkCommand == nil || linkCommand.Command == "" || len(linkCommand.Args) == 0 {
			return "", consts.ErrUnexpectedLinkValue
		}
		if err := os.WriteFile(
			linkCommand.Command,
			[]byte(stdout),
			fs.FileMode(os.O_CREATE),
		); err != nil {
			return "", err
		}
		return "", stderr
	default:
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
