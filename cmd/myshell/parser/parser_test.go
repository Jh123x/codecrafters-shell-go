package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFromReader(t *testing.T) {
	tests := map[string]struct {
		input string

		expectedCommand string
		expectedArgs    []string
		expectedErr     error
	}{
		"test inputs": {
			input:           "test arg1 arg2\n",
			expectedCommand: "test",
			expectedArgs:    []string{"arg1", "arg2"},
			expectedErr:     nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			cmd, arg, err := ParseFromReader(reader)
			assert.Equal(t, tc.expectedCommand, cmd)
			assert.Equal(t, tc.expectedArgs, arg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
