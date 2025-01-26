package parser

import (
	"bufio"
	"fmt"
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
	fmt.Println(command)
	split_args := strings.SplitN(
		strings.Trim(command, "\n"),
		" ", 2,
	)
	fmt.Println(split_args)

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
	for i := 0; i < len(argument); i++ {
		currentByte := argument[i]
		fmt.Println(currentByte)
		if currentByte == ' ' && currQuote == 0 {
			argStr = append(argStr, currArg.String())
			currArg.Reset()
			continue
		}

		if currentByte == currQuote {
			currQuote = 0
			argStr = append(argStr, currArg.String())
			currArg.Reset()
			continue
		}

		if currentByte == "'"[0] {
			currQuote = currentByte
			continue
		}

		currArg.WriteByte(currentByte)
	}

	if currArg.Len() > 0 {
		argStr = append(argStr, currArg.String())
	}

	fmt.Println(argStr)
	return argStr
}
