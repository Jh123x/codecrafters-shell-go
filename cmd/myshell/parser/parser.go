package parser

import (
	"bufio"
	"io"
	"strings"
)

func ParseFromReader(reader io.Reader) (string, []string, error) {
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		return "", nil, scanner.Err()
	}

	return parseCommand(scanner.Text())
}

func parseCommand(command string) (string, []string, error) {
	args, err := parseArguments(command)
	return args[0], args[1:], err
}

func parseArguments(argument string) ([]string, error) {
	currArg := strings.Builder{}
	argStr := make([]string, 0)
	currIdx := 0
	isEscape := false
	for currIdx < len(argument) {
		if isEscape {
			isEscape = false
			currArg.WriteByte(argument[currIdx])
			currIdx += 1
			continue
		}

		switch currentByte := argument[currIdx]; currentByte {
		case ' ':
			if currArg.Len() == 0 {
				break
			}

			argStr = append(argStr, currArg.String())
			currArg.Reset()
		case '\'', '"':
			arg, nextIdx, err := parseQuote(argument, currIdx)
			if err != nil {
				return nil, err
			}

			currArg.WriteString(arg)
			currIdx = nextIdx
			continue
		case '\\':
			isEscape = true
		default:
			currArg.WriteByte(currentByte)
		}
		currIdx += 1
	}

	if currArg.Len() > 0 {
		argStr = append(argStr, currArg.String())
	}

	return argStr, nil
}
