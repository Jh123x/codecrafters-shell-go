package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	tests := map[string]struct {
		args           []string
		expectedOutput string
		expectedErr    error
	}{
		"no args": {
			args:           []string{},
			expectedOutput: "\n",
		},
		"single arg": {
			args:           []string{"hello"},
			expectedOutput: "hello\n",
		},
		"multiple args": {
			args:           []string{"hello", "world!"},
			expectedOutput: "hello world!\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			output, err := Echo(tc.args)
			assert.Equal(t, tc.expectedOutput, output)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
