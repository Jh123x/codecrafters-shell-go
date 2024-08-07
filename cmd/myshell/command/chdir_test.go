package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/stretchr/testify/assert"
)

func TestChdir(t *testing.T) {
	oldDir, _ := os.Getwd()
	homeDir := filepath.Join(oldDir, "..", "command")
	os.Setenv(consts.ENV_HOME, homeDir)

	tests := map[string]struct {
		args        []string
		expectedErr error
		expectedDir string
	}{
		"no args": {
			args:        []string{},
			expectedDir: oldDir,
		},
		"unknown dir": {
			args:        []string{"/unknown"},
			expectedErr: fmt.Errorf("cd: /unknown: No such file or directory"),
			expectedDir: oldDir,
		},
		"parent dir": {
			args:        []string{".."},
			expectedDir: strings.TrimSuffix(oldDir, "command")[:len(oldDir)-8],
		},
		"current dir": {
			args:        []string{"."},
			expectedDir: oldDir,
		},
		"home dir": {
			args:        []string{"~"},
			expectedDir: homeDir,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer os.Chdir(oldDir)
			_, err := ChangeDir(tc.args)
			assert.Equal(t, tc.expectedErr, err)
			cwd, _ := os.Getwd()
			assert.Equal(t, tc.expectedDir, cwd)
		})
	}
}
