package command

import (
	"errors"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/parser"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	tests := map[string]struct {
		command *parser.Command

		expectedOutput string
		expectedError  error
	}{
		"not found": {
			command: parser.NewCommand(
				"test_not_found",
				[]string{"arg1", "arg2"},
			),
			expectedOutput: "",
			expectedError:  errors.New("test_not_found: command not found"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := DefaultCommand(tc.command)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedOutput, output, name)
		})
	}
}
