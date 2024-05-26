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
	split_args := strings.Split(
		strings.Trim(command, "\n"),
		" ",
	)

	return split_args[0], split_args[1:]
}
