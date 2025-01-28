package command

import (
	"errors"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/files"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func DefaultCommand(command *parser.Command) (string, error) {
	filePath, err := files.GetFilePath(command.Command)
	if err != nil {
		switch err {
		case consts.ErrFileNotFound:
			return parseNotfoundError(command.Command)
		default:
			return "", err
		}
	}

	return files.RunFile(filePath, command.Args)
}

func parseNotfoundError(command string) (string, error) {
	builder := strings.Builder{}
	builder.WriteString(command)
	builder.WriteString(": command not found")
	return "", errors.New(builder.String())
}
