package command

import (
	"io/fs"
	"os"
)

func WriteToFile(filePath, contents string) error {
	return os.WriteFile(filePath, []byte(contents), fs.ModePerm)
}

func AppendToFile(filePath, contents string) error {
	fd, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer fd.Close()

	if _, err := fd.WriteString(contents); err != nil {
		return err
	}

	return nil
}
