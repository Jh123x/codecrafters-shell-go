package consts

import "errors"

const (
	EXIT = "exit"
	ECHO = "echo"
)

var (
	TypeMap = map[string]string{
		EXIT: "exit is a shell builtin",
		ECHO: "echo is a shell builtin",
	}
)

var (
	ErrEXIT      = errors.New("exit")
	ErrTypeUsage = errors.New("usage: type <command>")
)
