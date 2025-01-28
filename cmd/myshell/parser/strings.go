package parser

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrIncompleteQuote     = errors.New("missing closing quote")
	ErrUnexpectedEndOfLine = errors.New("unexpected end of line")
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
	for startIdx < len(arguments) {
		fmt.Println(startIdx, string([]byte{arguments[startIdx]}))
		switch currByte := arguments[startIdx]; currByte {
		case '"':
			return builder.String(), startIdx + 1, nil
		case '\\':
			val, nextIdx, err := parseEscape(arguments, startIdx+1)
			if err != nil {
				return "", -1, err
			}
			builder.WriteString(val)
			startIdx = nextIdx
		default:
			builder.WriteByte(currByte)
		}

		startIdx += 1
	}

	return "", -1, ErrIncompleteQuote
}

func parseEscape(arguments string, startIdx int) (string, int, error) {
	if startIdx >= len(arguments) {
		return "", -1, ErrUnexpectedEndOfLine
	}

	switch currByte := arguments[startIdx]; currByte {
	case 'n':
		return "\n", startIdx, nil
	case '\\':
		return "\\", startIdx, nil
	case '$':
		return "$", startIdx, nil
	case '"':
		return "\"", startIdx, nil
	default:
		return string([]byte{'\\', currByte}), startIdx, nil
	}
}
