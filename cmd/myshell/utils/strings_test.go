package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixStrPrinting(t *testing.T) {
	tests := map[string]struct {
		val         string
		expectedRes string
	}{
		"printing should not touch \r\n": {
			val:         "this is my \r\n string\r\n",
			expectedRes: "this is my \r\n string\r\n",
		},
		"printing should fix \n": {
			val:         "this is my \n string\n",
			expectedRes: "this is my \r\n string\r\n",
		},
		"Double apply should have no effect": {
			val:         FixStrPrinting("this is my \n string\n"),
			expectedRes: "this is my \r\n string\r\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, FixStrPrinting(tc.val))
		})
	}
}
