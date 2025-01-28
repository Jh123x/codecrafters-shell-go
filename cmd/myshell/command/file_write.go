package command

import (
	"errors"
	"io/fs"
	"os"
)

func WriteToFile(filePath, contents string) error {
	return os.WriteFile(filePath, []byte(contents), fs.ModePerm)
}

func AppendToFile(filePath, contents string) error {
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return WriteToFile(filePath, contents)
		}
		return err
	}

	fd, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer fd.Close()

	if _, err := fd.WriteString(contents); err != nil {
		return err
	}

	return nil
}
