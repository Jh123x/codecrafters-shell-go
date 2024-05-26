package files

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func GetFilePath(filename string) (string, error) {
	envPaths := os.Getenv(consts.ENV_PATH)
	if len(envPaths) == 0 {
		return "", consts.ErrEnvPathNotSet
	}

	for _, envPath := range strings.Split(envPaths, ":") {
		filePath := path.Join(envPath, filename)
		if _, err := os.Stat(filePath); err != nil {
			continue
		}

		return filePath, nil
	}
	return "", consts.ErrFileNotFound
}

func RunFile(absFilePath string, args []string) (string, error) {
	cmd := exec.Command(absFilePath, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
