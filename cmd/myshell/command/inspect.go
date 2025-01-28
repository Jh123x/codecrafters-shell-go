package command

import (
	"encoding/json"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
)

func Inspect(cmd *parser.Command) (string, error) {
	res, err := json.Marshal(cmd)
	if err != nil {
		return "", err
	}

	cmd.Link = nil
	return string(append(res, '\n')), nil
}
