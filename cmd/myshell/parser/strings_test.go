package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseQuote(t *testing.T) {
	tests := map[string]struct {
		arg      string
		startIdx int

		expectedRes     string
		expectedNextIdx int
		expectedErr     error
	}{
		"proper single quote": {
			arg:             `'test'`,
			startIdx:        0,
			expectedRes:     "test",
			expectedNextIdx: 6,
			expectedErr:     nil,
		},
		"proper single quote example": {
			arg:             `'shell\\\nscript'`,
			startIdx:        0,
			expectedRes:     `shell\\\nscript`,
			expectedNextIdx: 17,
			expectedErr:     nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			arg, end, err := parseQuote(tc.arg, tc.startIdx)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedNextIdx, end)
			assert.Equal(t, tc.expectedRes, arg)
		})
	}
}
