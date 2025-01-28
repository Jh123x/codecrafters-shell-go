package parser

import (
	"errors"
	"strings"
)

var (
	ErrIncompleteQuote = errors.New("missing closing quote")
)

func parseQuote(arguments string, startIdx int) (string, int, error) {
	quote := arguments[startIdx]

	switch quote {
	case '"':
		return parseDoubleQuote(arguments, startIdx+1)
	case '\'':
		return parseSingleQuote(arguments, startIdx+1)
	default:
		return "", -1, nil
	}
}

func parseSingleQuote(arguments string, startIdx int) (string, int, error) {
	builder := strings.Builder{}

	for startIdx < len(arguments) {
		switch currByte := arguments[startIdx]; currByte {
		case '\'':
			return builder.String(), startIdx + 1, nil
		default:
			builder.WriteByte(currByte)
		}
		startIdx += 1
	}

	return "", -1, ErrIncompleteQuote
}

func parseDoubleQuote(arguments string, startIdx int) (string, int, error) {
	builder := strings.Builder{}
	isEscape := false
	for startIdx < len(arguments) {
		if isEscape {
			isEscape = false
			builder.WriteByte(arguments[startIdx])
			startIdx += 1
			continue
		}

		switch currByte := arguments[startIdx]; currByte {
		case '"':
			return builder.String(), startIdx + 1, nil
		case '\\':
			isEscape = true
		default:
			builder.WriteByte(currByte)
		}

		startIdx += 1
	}

	return "", -1, ErrIncompleteQuote
}
