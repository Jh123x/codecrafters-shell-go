package command

import (
	"errors"
	"strings"
)

func NotFound(command string) (string, error) {
	builder := strings.Builder{}
	builder.WriteString(command)
	builder.WriteString(": command not found")
	return "", errors.New(builder.String())
}
