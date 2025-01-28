package consts

import "errors"

const (
	// Commands
	EXIT    = "exit"
	ECHO    = "echo"
	TYPE    = "type"
	PWD     = "pwd"
	CD      = "cd"
	Inspect = "inspect"

	// Environment variables
	ENV_PATH = "PATH"
	ENV_HOME = "HOME"
)

type empty struct{}

var (
	TypeMap = map[string]empty{
		EXIT: {},
		ECHO: {},
		TYPE: {},
		PWD:  {},
		CD:   {},
	}
)

var (
	ErrEXIT          = errors.New("exit")
	ErrEnvPathNotSet = errors.New("environment variable PATH not set")
	ErrFileNotFound  = errors.New("file not found")

	ErrTypeUsage  = errors.New("usage: type <command>")
	ErrChdirUsage = errors.New("usage: cd <directory>")

	ErrEmptyCommand        = errors.New("command cannot be empty")
	ErrIncompleteQuote     = errors.New("missing closing quote")
	ErrUnexpectedEndOfLine = errors.New("unexpected end of line")
	ErrInvalidCommandStart = errors.New("invalid command start")
	ErrUnsupportedLinkType = errors.New("this link type is not supported yet")
	ErrNilLink             = errors.New("link cannot be nil")
	ErrUnexpectedLinkValue = errors.New("unexpected link error")
)
