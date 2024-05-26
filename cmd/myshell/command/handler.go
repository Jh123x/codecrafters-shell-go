package command

func HandleCommand(command string, args []string) (string, error) {
	return NotFound(command)
}
