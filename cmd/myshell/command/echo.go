package command

import "strings"

func Echo(args []string) (string, error) {
	return strings.Join(args, " ") + "\n", nil
}
