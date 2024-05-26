package consts

import "errors"

const (
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
)

var (
	TypeMap = map[string]string{
		EXIT: "exit is a shell builtin",
		ECHO: "echo is a shell builtin",
		TYPE: "type is a shell builtin",
	}
)

var (
	ErrEXIT      = errors.New("exit")
	ErrTypeUsage = errors.New("usage: type <command>")
)
