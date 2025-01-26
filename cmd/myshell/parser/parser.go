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

	cmd, args := parseCommand(scanner.Text())
	return cmd, args, nil
}

func parseCommand(command string) (string, []string) {
	split_args := strings.SplitN(
		strings.TrimRight(command, "\n"),
		" ", 2,
	)

	switch len(split_args) {
	case 0:
		return "", nil
	case 1:
		return split_args[0], nil
	default:
		return split_args[0], parseArguments(split_args[1])
	}
}

func parseArguments(argument string) []string {
	currArg := strings.Builder{}
	argStr := make([]string, 0)
	currQuote := byte(0)
	isEscape := false
	for i := 0; i < len(argument); i++ {
		currentByte := argument[i]
		if isEscape && currQuote == byte(0) {
			isEscape = false
			currArg.WriteByte(currentByte)
			continue
		}

		if currentByte == ' ' && currQuote == 0 {
			if currArg.Len() > 0 {
				argStr = append(argStr, currArg.String())
				currArg.Reset()
			}

			continue
		}

		if currentByte == '\\' && currQuote == byte(0) {
			isEscape = true
			continue
		}

		if currentByte == currQuote {
			currQuote = 0
			continue
		}

		if currQuote == byte(0) && currentByte == '\'' || currentByte == '"' {
			currQuote = currentByte
			continue
		}

		currArg.WriteByte(currentByte)
	}

	if currArg.Len() > 0 {
		argStr = append(argStr, currArg.String())
	}

	return argStr
}
