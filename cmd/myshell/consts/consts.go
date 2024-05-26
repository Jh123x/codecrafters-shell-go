package consts

import "errors"

const (
	// Commands
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"

	// Environment variables
	ENV_PATH = "PATH"
)

var (
	TypeMap = map[string]string{
		EXIT: "exit is a shell builtin",
		ECHO: "echo is a shell builtin",
		TYPE: "type is a shell builtin",
	}
)

var (
	ErrEXIT          = errors.New("exit")
	ErrTypeUsage     = errors.New("usage: type <command>")
	ErrEnvPathNotSet = errors.New("environment variable PATH not set")
	ErrFileNotFound  = errors.New("file not found")
)
