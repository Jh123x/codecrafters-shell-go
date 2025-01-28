package autocomplete

import (
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

func GetClosestCommands(currBuffer string) ([]string, error) {
	foundCommands := make(map[string]struct{}, 10)
	// Search from builtins.
	for fn := range consts.TypeMap {
		if len(fn) >= len(currBuffer) && strings.HasPrefix(fn, currBuffer) {
			foundCommands[fn] = struct{}{}
		}
	}

	// Search from path env
	envPaths := os.Getenv(consts.ENV_PATH)
	if len(envPaths) == 0 {
		return getKeys(foundCommands), nil
	}

	for _, envPath := range strings.Split(envPaths, ":") {
		dir, err := os.ReadDir(envPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}

		for _, dirEntry := range dir {
			entryName := dirEntry.Name()
			if !dirEntry.IsDir() && len(entryName) >= len(currBuffer) && strings.HasPrefix(entryName, currBuffer) {
				foundCommands[entryName] = struct{}{}
			}
		}
	}

	return getKeys(foundCommands), nil
}

func getKeys(v map[string]struct{}) []string {
	acc := make([]string, 0, len(v))

	for name := range v {
		acc = append(acc, name)
	}

	return acc
}

func GetCommonPrefix(commands []string) string {
	currPrefix := commands[0]
	for _, cmd := range commands[1:] {
		sharedChars := getSharedChars(currPrefix, cmd)
		currPrefix = currPrefix[:sharedChars]
	}
	return currPrefix
}

func getSharedChars(s1, s2 string) int {
	minLen := min(len(s1), len(s2))
	for idx := 0; idx < minLen; idx++ {
		if s1[idx] == s2[idx] {
			continue
		}
		return idx
	}
	return minLen
}
