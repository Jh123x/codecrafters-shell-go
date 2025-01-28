package command

import (
	"encoding/json"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

// Inspect is used for shell debugging.
func Inspect(cmd *parser.Command) (string, error) {
	res, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	builder.WriteString("Raw Struct:\n")
	builder.Write(res)
	builder.WriteString("\nLink Info:\n")
	if cmd.Link != nil {
		builder.WriteString(cmd.Link.GetInfo())
	} else {
		builder.WriteString("no link")
	}

	cmd.Link = nil

	builder.WriteString("\n")
	return builder.String(), nil
}
