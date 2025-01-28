package autocomplete

import (
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func GetClosestCommand(currBuffer string) (string, error) {
	for fn := range consts.TypeMap {
		if strings.Contains(fn, currBuffer) {
			return fn, nil
		}
	}

	return "", nil
}
