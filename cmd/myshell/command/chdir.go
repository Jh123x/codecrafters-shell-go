package command

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func ChangeDir(args []string) (string, error) {
	if len(args) != 1 {
		return "", consts.ErrChdirUsage
	}

	dir := args[0]
	if dir == "~" {
		dir = os.Getenv(consts.ENV_HOME)
	}

	if err := os.Chdir(dir); err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory", dir)
	}

	return "", nil
}
