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
			expectedOutput: "",
		},
		"single arg": {
			args:           []string{"hello"},
			expectedOutput: "hello",
		},
		"multiple args": {
			args:           []string{"hello", "world!"},
			expectedOutput: "hello world!",
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
