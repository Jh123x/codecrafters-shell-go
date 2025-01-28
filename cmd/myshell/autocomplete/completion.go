package autocomplete

import (
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func GetClosestCommand(currBuffer string) (string, error) {
	// Search from builtins.
	for fn := range consts.TypeMap {
		if strings.Contains(fn, currBuffer) {
			return fn, nil
		}
	}

	// Search from path env
	envPaths := os.Getenv(consts.ENV_PATH)
	if len(envPaths) == 0 {
		return "", nil
	}

	for _, envPath := range strings.Split(envPaths, ":") {
		dir, err := os.ReadDir(envPath)
		if err != nil {
			return "", err
		}

		for _, dirEntry := range dir {
			if !strings.Contains(dirEntry.Name(), currBuffer) {
				continue
			}
			return dirEntry.Name(), nil
		}
	}

	return "", nil
}
