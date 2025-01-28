package parser

type LinkType int

const (
	LinkTypeStdout LinkType = iota
	LinkTypeStderr
	LinkTypeAppendStdout
	LinkTypeAppendStderr
)

type Link struct {
	Type          LinkType
	LinkedCommand *Command
}

type Command struct {
	Command string
	Args    []string
	Next    *Link
}

func NewCommand(
	command string,
	args []string,
) *Command {
	return &Command{
		Command: command,
		Args:    args,
		Next:    nil,
	}
}

func (c *Command) Link(link *Link) { c.Next = link }
