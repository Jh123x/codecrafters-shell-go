package command

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func Type(args []string) (string, error) {
	if len(args) != 1 {
		return "", consts.ErrTypeUsage
	}

	commandType := args[0]
	helpMsg, ok := consts.TypeMap[commandType]
	if !ok {
		return "", fmt.Errorf("%s not found", commandType)
	}

	return helpMsg, nil
}
