package parser

import (
	"fmt"
	"strings"
)

type LinkType int

func (lt LinkType) String() string {
	switch lt {
	case LinkTypeStdout:
		return "Redirect stdout to file (write)"
	case LinkTypeStderr:
		return "Redirect stderr to file (write)"
	case LinkTypePipe:
		return "Pipe"
	case LinkTypeAppendStderr:
		return "Redirect stderr to file (append)"
	case LinkTypeAppendStdout:
		return "Redirect stdout to file (append)"
	default:
		return fmt.Sprintf("%d", lt)
	}
}

const (
	LinkTypeStdout LinkType = iota
	LinkTypeStderr
	LinkTypeAppendStdout
	LinkTypeAppendStderr

	// Others
	LinkTypePipe
	LinkTypeNone
)

type Link struct {
	Type LinkType
	Args []string
}

func (l *Link) GetInfo() string {
	builder := strings.Builder{}
	builder.WriteString(l.Type.String())
	builder.WriteString("\nWith Args: ")
	builder.WriteString(strings.Join(l.Args, ","))
	return builder.String()
}

type Command struct {
	Command string
	Args    []string
	Link    *Link
}

func NewCommand(
	command string,
	args []string,
) *Command {
	return &Command{
		Command: command,
		Args:    args,
		Link:    nil,
	}
}
