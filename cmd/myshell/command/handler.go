package command

import "io"

func HandleCommand(writer io.Writer, command string, args []string) error {
	switch command {
	default:
		return NotFound(writer, command)
	}
}
