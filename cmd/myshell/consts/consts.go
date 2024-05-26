package consts

import "errors"

const (
	EXIT = "exit"
	ECHO = "echo"
)

var (
	ErrEXIT = errors.New("exit")
)
