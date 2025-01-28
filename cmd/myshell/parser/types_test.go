package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	defer func() { assert.Nil(t, recover()) }()
	res := NewCommand("test", []string{"test1", "test2"})
	assert.Equal(
		t,
		&Command{
			Command: "test",
			Args:    []string{"test1", "test2"},
		},
		res,
	)
}
