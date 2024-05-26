package command

import (
	"errors"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	tests := map[string]struct {
		args []string

		expectedOutput string
		expectedErr    error
	}{
		"no args": {
			args:           []string{},
			expectedOutput: "",
			expectedErr:    consts.ErrTypeUsage,
		},
		"builtin command": {
			args:           []string{consts.EXIT},
			expectedOutput: "exit is a shell builtin\n",
		},
		"unknown command": {
			args:           []string{"unknown"},
			expectedOutput: "",
			expectedErr:    errors.New("unknown not found\n"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := Type(tc.args)
			assert.Equal(t, tc.expectedOutput, output)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
