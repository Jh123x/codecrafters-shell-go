package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFromReader(t *testing.T) {
	tests := map[string]struct {
		input string

		expectedCommand *Command
		expectedErr     error
	}{
		"test inputs": {
			input:           "test arg1 arg2\n",
			expectedCommand: NewCommand("test", []string{"arg1", "arg2"}),
			expectedErr:     nil,
		},
		"test with quotes": {
			input:           `'exe with "quotes"' file`,
			expectedCommand: NewCommand(`exe with "quotes"`, []string{"file"}),
			expectedErr:     nil,
		},
		"test with write to out file": {
			input: `cat test.txt > test.txt`,
			expectedCommand: &Command{
				Command: "cat",
				Args:    []string{"test.txt"},
				Link: &Link{
					Type: LinkTypeStdout,
					Args: []string{"test.txt"},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			cmd, err := ParseFromReader(reader)
			assert.Equal(t, tc.expectedCommand, cmd)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func Test_parseArguments(t *testing.T) {
	tests := map[string]struct {
		argStr       string
		expectedArgs []string
		expectedErr  error
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
		"double quotes": {
			argStr:       `test "test2""test3"`,
			expectedArgs: []string{"test", "test2test3"},
		},
		"double quotes case": {
			argStr:       `"bar"  "shell's"  "foo"`,
			expectedArgs: []string{"bar", "shell's", "foo"},
		},
		"sample 3": {
			argStr: `world\ \ \ \ \ \ script`,
			expectedArgs: []string{
				"world      script",
			},
		},
		"single quotes eg": {
			argStr: `'shell\\\nscript'`,
			expectedArgs: []string{
				"shell\\\\\\nscript",
			},
		},
		"single quotes eg2": {
			argStr: "'example\\\"testhello\\\"shell'",
			expectedArgs: []string{
				"example\\\"testhello\\\"shell",
			},
		},
		"quote test": {
			argStr: "\"/tmp/baz/'f 81'\" \"/tmp/baz/'f  \\58'\" \"/tmp/baz/'f \\83\\'\"",
			expectedArgs: []string{
				`/tmp/baz/'f 81'`,
				`/tmp/baz/'f  \58'`,
				`/tmp/baz/'f \83\'`,
			},
		},
		"quote test 2": {
			argStr: "\"/tmp/baz/f\\n9\" \"/tmp/baz/f\\42\" \"/tmp/baz/f'\\'62\"",
			expectedArgs: []string{
				"/tmp/baz/f\\n9",
				"/tmp/baz/f\\42",
				"/tmp/baz/f'\\'62",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			args, err := parseArguments(tc.argStr)
			assert.Equal(t, tc.expectedArgs, args)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
