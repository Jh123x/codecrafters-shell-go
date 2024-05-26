package command

import (
	"fmt"
	"io"
)

func NotFound(reader io.Writer, command string) error {
	_, err := fmt.Fprintf(reader, "%s: command not found", command)
	if err != nil {
		return err
	}

	return nil
}
