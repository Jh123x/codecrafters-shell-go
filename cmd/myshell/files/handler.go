package files

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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
	absFilePath = filepath.Clean(absFilePath)
	cmd := exec.Command(absFilePath, args...)
	cmd.Args[0] = filepath.Base(absFilePath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	outB, err := io.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	errB, err := io.ReadAll(stderr)
	if err != nil {
		return "", err
	}

	return string(outB), parseErrStr(string(errB))
}

func parseErrStr(errorVal string) error {
	if len(errorVal) == 0 {
		return nil
	}

	return errors.New(errorVal)
}
