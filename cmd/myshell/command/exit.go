package command

import "github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"

func Exit(args []string) (string, error) {
	return "", consts.ErrEXIT
}
