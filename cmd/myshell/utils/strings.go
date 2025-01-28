package utils

import "strings"

func FixStrPrinting(val string) string {
	res := strings.ReplaceAll(
		strings.ReplaceAll(val, "\r\n", "\n"),
		"\n",
		"\r\n",
	)

	if !strings.HasSuffix(res, "\r\n") {
		return res + "\r\n"
	}

	return res
}
