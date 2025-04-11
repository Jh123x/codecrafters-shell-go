package autocomplete

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/trie"
)

var autoComplete *trie.Trie

func init() {
	if autoComplete != nil {
		return
	}
	autoComplete = trie.NewTrie()

	for v := range consts.TypeMap {
		autoComplete.AddWords(v)
	}

	// Search from path env
	envPaths := os.Getenv(consts.ENV_PATH)
	for _, envPath := range strings.Split(envPaths, ":") {
		dir, err := os.ReadDir(envPath)
		if err != nil {
			return
		}

		for _, dirEntry := range dir {
			entryName := dirEntry.Name()
			if strings.Contains(entryName, ".") {
				continue
			}
			if dirEntry.IsDir() {
				continue
			}

			absPath := filepath.Join(envPath, dirEntry.Name())
			if _, err := os.Stat(absPath); err != nil {
				fmt.Println("skipped", absPath, err.Error())
				continue
			}

			autoComplete.AddWords(entryName)
		}
	}
}

func GetClosestCommands(currBuffer string) ([]string, error) {
	return autoComplete.GetCompletion(currBuffer), nil
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
