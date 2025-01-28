package command

import (
	"fmt"
	"os"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	cwd, _ := os.Getwd()

	tests := map[string]struct {
		command string
		args    []string

		expectedOutput string
		expectedErr    error
	}{
		"not found": {
			command:        "test_not_found",
			args:           []string{"should", "be", "ignored"},
			expectedOutput: "",
			expectedErr:    fmt.Errorf("test_not_found: command not found\n"),
		},
		"echo": {
			command:        consts.ECHO,
			args:           []string{"hello", "world"},
			expectedOutput: "hello world\n",
		},
		"exit": {
			command:        consts.EXIT,
			args:           []string{"0"},
			expectedOutput: "",
			expectedErr:    consts.ErrEXIT,
		},
		"type": {
			command:        consts.TYPE,
			args:           []string{"exit"},
			expectedOutput: "exit is a shell builtin\n",
		},
		"type no args": {
			command:        consts.TYPE,
			args:           []string{},
			expectedOutput: "",
			expectedErr:    consts.ErrTypeUsage,
		},
		"pwd": {
			command:        consts.PWD,
			args:           []string{},
			expectedOutput: cwd + "\n",
		},
		"cd": {
			command: consts.CD,
			args:    []string{"."},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := HandleCommand(parser.NewCommand(tc.command, tc.args))

			assert.Equal(t, tc.expectedOutput, output)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
