package autocomplete

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommonPrefix(t *testing.T) {
	tests := map[string]struct {
		commands             []string
		expectedCommonPrefix string
	}{
		"test with common": {
			commands:             []string{"test", "test1", "test2", "test3"},
			expectedCommonPrefix: "test",
		},
		"no common": {
			commands:             []string{"test", "not a test"},
			expectedCommonPrefix: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCommonPrefix, GetCommonPrefix(tc.commands))
		})
	}
}
