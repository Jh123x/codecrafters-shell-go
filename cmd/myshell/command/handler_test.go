package command

import (
	"fmt"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := map[string]struct {
		command string
		args    []string

		expectedOutput string
		expectedErr    error
	}{
		"not found": {
			command:        "test",
			args:           []string{"should", "be", "ignored"},
			expectedOutput: "",
			expectedErr:    fmt.Errorf("test: command not found"),
		},
		"echo": {
			command:        consts.ECHO,
			args:           []string{"hello", "world"},
			expectedOutput: "hello world",
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
			expectedOutput: "exit is a shell builtin",
		},
		"type no args": {
			command:        consts.TYPE,
			args:           []string{},
			expectedOutput: "",
			expectedErr:    consts.ErrTypeUsage,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := HandleCommand(tc.command, tc.args)

			assert.Equal(t, tc.expectedOutput, output)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
