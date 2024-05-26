package command

import (
	"fmt"
	"testing"

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
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := HandleCommand(tc.command, tc.args)

			assert.Equal(t, tc.expectedOutput, output)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
