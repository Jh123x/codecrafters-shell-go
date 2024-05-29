package command

import (
	"fmt"
	"os"
)

func Pwd() (string, error) {
	res, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", res), nil
}
