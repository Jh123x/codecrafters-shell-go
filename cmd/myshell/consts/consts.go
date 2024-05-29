package consts

import "errors"

const (
	// Commands
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
	PWD  = "pwd"

	// Environment variables
	ENV_PATH = "PATH"
)

var (
	TypeMap = map[string]string{
		EXIT: "exit is a shell builtin\n",
		ECHO: "echo is a shell builtin\n",
		TYPE: "type is a shell builtin\n",
	}
)

var (
	ErrEXIT          = errors.New("exit")
	ErrTypeUsage     = errors.New("usage: type <command>")
	ErrEnvPathNotSet = errors.New("environment variable PATH not set")
	ErrFileNotFound  = errors.New("file not found")
)
