package command

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func ChangeDir(args []string) (string, error) {
	if len(args) > 1 {
		return "", consts.ErrChdirUsage
	}

	dir := ""
	if len(args) == 0 || args[0] == "~" {
		dir = os.Getenv(consts.ENV_HOME)
	} else {
		dir = args[0]
	}

	return "", changeDir(dir)
}

func changeDir(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory\n", dir)
	}

	return nil
}
