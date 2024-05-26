package command

import "github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"

func HandleCommand(command string, args []string) (string, error) {
	switch command {
	case consts.ECHO:
		return Echo(args)
	case consts.EXIT:
		return Exit(args)
	default:
		return NotFound(command)
	}
}
