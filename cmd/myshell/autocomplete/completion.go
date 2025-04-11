package autocomplete

import (
	"os"
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
			if dirEntry.IsDir() {
				continue
			}

			if stat, err := os.Stat(dirEntry.Name()); err != nil || stat.Mode().Perm()&0111 == 0 {
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
