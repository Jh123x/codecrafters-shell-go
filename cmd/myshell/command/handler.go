package command

import (
	"errors"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

func HandleCommand(command *parser.Command) (string, error) {
	if command == nil {
		return "", consts.ErrEmptyCommand
	}

	stdout, err := handleCommand(command)
	if len(stdout) > 0 {
		stdout = utils.FixStrPrinting(stdout)
	}

	if err != nil && err != consts.ErrEXIT {
		err = errors.New(utils.FixStrPrinting(err.Error()))
	}

	if command.Link != nil {
		return handleLink(command.Link, stdout, err)
	}

	return stdout, err
}

func handleLink(link *parser.Link, stdout string, stderr error) (string, error) {
	if link == nil || len(link.Args) != 1 {
		return "", consts.ErrUnexpectedLinkValue
	}

	fileName := link.Args[0]
	switch link.Type {
	case parser.LinkTypeStdout:
		if err := WriteToFile(fileName, stdout); err != nil {
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

		if err := WriteToFile(fileName, errTxt); err != nil {
			return "", err
		}

		return stdout, nil
	case parser.LinkTypeAppendStdout:
		if err := AppendToFile(fileName, stdout); err != nil {
			return "", err
		}
		if stderr == nil {
			return "", nil
		}
		return "", stderr
	case parser.LinkTypeAppendStderr:
		errTxt := ""
		if stderr != nil {
			errTxt = stderr.Error()
		}

		if err := AppendToFile(fileName, errTxt); err != nil {
			return "", err
		}

		return stdout, nil
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

	// For debugging command struct.
	case consts.INSPECT:
		return Inspect(command)

	// Executable files.
	default:
		return DefaultCommand(command)
	}
}
