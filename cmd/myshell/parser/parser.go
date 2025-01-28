package parser

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func ParseCommands(command string) (*Command, error) {
	args, err := parseArguments(command)
	if err != nil {
		return nil, err
	}

	return parseCommand(args)
}

func parseCommand(tokens []string) (*Command, error) {
	if len(tokens) == 0 {
		return nil, nil
	}

	var (
		currIdx int = 0
		cmd     *Command
	)

	for currIdx < len(tokens) {
		switch curr := tokens[currIdx]; curr {
		case ">", ">>", "2>>", "2>", "1>", "1>>":
			if cmd == nil {
				return nil, consts.ErrInvalidCommandStart
			}

			link, err := parseLink(curr, tokens[currIdx+1:])
			if err != nil {
				return nil, err
			}

			cmd.Link = link
			return cmd, nil
		default:
			if cmd == nil {
				cmd = &Command{}
			}

			if cmd.Command == "" {
				cmd.Command = curr
				break
			}

			cmd.Args = append(cmd.Args, curr)
		}
		currIdx += 1
	}

	return cmd, nil
}

func parseLink(linkType string, tokens []string) (*Link, error) {
	var linkVal LinkType
	switch linkType {
	case ">", "1>":
		linkVal = LinkTypeStdout
	case "2>":
		linkVal = LinkTypeStderr
	case ">>", "1>>":
		linkVal = LinkTypeAppendStdout
	case "2>>":
		linkVal = LinkTypeAppendStderr
	case "|":
		linkVal = LinkTypePipe
	case ";":
		linkVal = LinkTypeNone
	default:
		return nil, consts.ErrUnsupportedLinkType
	}

	return &Link{Type: linkVal, Args: tokens}, nil
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
