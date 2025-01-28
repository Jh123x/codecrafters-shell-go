package autocomplete

import "github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"

func GetClosestCommand(currBuffer string) (string, error) {
	switch currBuffer {
	case "exi":
		return consts.EXIT, nil
	case "ech":
		return consts.ECHO, nil
	default:
		return "", nil
	}
}
