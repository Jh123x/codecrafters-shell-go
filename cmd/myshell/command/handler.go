package command

import "github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"

func HandleCommand(command string, args []string) (string, error) {
	switch command {
	case consts.ECHO:
		return Echo(args)
	case consts.EXIT:
		return Exit(args)
	case consts.TYPE:
		return Type(args)
	case consts.PWD:
		return Pwd()
	case consts.CD:
		return ChangeDir(args)
	default:
		return DefaultCommand(command, args)
	}
}
