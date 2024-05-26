package command

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/files"
)

func Type(args []string) (string, error) {
	if len(args) != 1 {
		return "", consts.ErrTypeUsage
	}

	typeArg := args[0]
	helpMsg, ok := consts.TypeMap[typeArg]
	if ok {
		return helpMsg, nil
	}

	if absPath, err := files.GetFilePath(typeArg); err == nil {
		return fmt.Sprintf("%s is %s\n", typeArg, absPath), nil
	}

	return "", fmt.Errorf("%s not found", typeArg)
}
