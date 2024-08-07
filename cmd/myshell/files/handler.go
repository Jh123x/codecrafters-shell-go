package files

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func GetFilePath(filename string) (string, error) {
	if path.IsAbs(filename) {
		return filename, nil
	}

	return ParseRelativePath(filename)
}

func ParseRelativePath(filename string) (string, error) {
	// Check if the file is in the current directory
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

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
