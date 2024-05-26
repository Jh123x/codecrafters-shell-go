package command

import (
	"errors"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/files"
)

func DefaultCommand(command string, args []string) (string, error) {
	filePath, err := files.GetFilePath(command)
	if err != nil {
		switch err {
		case consts.ErrFileNotFound:
			return parseNotfoundError(command)
		default:
			return "", err
		}
	}

	return files.RunFile(filePath, args)
}

func parseNotfoundError(command string) (string, error) {
	builder := strings.Builder{}
	builder.WriteString(command)
	builder.WriteString(": command not found")
	return "", errors.New(builder.String())
}
