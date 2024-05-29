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

type empty struct{}

var (
	TypeMap = map[string]empty{
		EXIT: {},
		ECHO: {},
		TYPE: {},
		PWD:  {},
	}
)

var (
	ErrEXIT          = errors.New("exit")
	ErrTypeUsage     = errors.New("usage: type <command>")
	ErrEnvPathNotSet = errors.New("environment variable PATH not set")
	ErrFileNotFound  = errors.New("file not found")
)
