package command

import (
	"io/fs"
	"os"
)

func WriteToFile(filePath, contents string) error {
	return os.WriteFile(filePath, []byte(contents), fs.ModePerm)
}
