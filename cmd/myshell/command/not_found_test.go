package command

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	output, err := NotFound("test")
	assert.Equal(t, "", output)
	assert.Equal(t, errors.New("test: command not found"), err)
}
