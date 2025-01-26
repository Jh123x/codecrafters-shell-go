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

func Test_parseArguments(t *testing.T) {
	tests := map[string]struct {
		argStr       string
		expectedArgs []string
	}{
		"no args": {
			argStr:       "",
			expectedArgs: []string{},
		},
		"single arg": {
			argStr:       "this_is_one_arg",
			expectedArgs: []string{"this_is_one_arg"},
		},
		"extra spaces": {
			argStr:       "test      test2",
			expectedArgs: []string{"test", "test2"},
		},
		"test special quotes case": {
			argStr:       "'test     example' 'hello''world'",
			expectedArgs: []string{"test     example", "helloworld"},
		},
		"no single quotes": {
			argStr:       "test test",
			expectedArgs: []string{"test", "test"},
		},
		"single quotes": {
			argStr:       "test 'test test'",
			expectedArgs: []string{"test", "test test"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedArgs, parseArguments(tc.argStr))
		})
	}
}
