package parser

type LinkType int

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
	Type          LinkType
	LinkedCommand *Command
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
