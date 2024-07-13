package command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	tests := map[string]struct {
		command string
		args    []string

		expectedOutput string
		expectedError  error
	}{
		"not found": {
			command: "test_not_found",
			args: []string{
				"arg1",
				"arg2",
			},
			expectedOutput: "",
			expectedError:  errors.New("test_not_found: command not found"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := DefaultCommand(tc.command, tc.args)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedOutput, output, name)
		})
	}
}
