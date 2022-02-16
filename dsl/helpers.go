package dsl

import (
	"errors"
	"os"
	"strings"
)

func getKeyValue(str string) (string, []string) {
	c := strings.Split(str, " ")
	if len(c) == 0 {
		return "", nil
	}
	return strings.TrimSpace(c[0]), c[1:]
}
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func isFileExists(name string) bool {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func clearWhiteSpace(s []string) (result []string) {
	for _, i := range s {
		if len(i) > 0 {
			result = append(result, i)
		}
	}
	return result
}
