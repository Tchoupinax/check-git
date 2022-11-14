package utils

import (
	"fmt"

	"strings"
)

func AddSpaceToEnd(s string, finalExpectedSize int) string {
	diff := finalExpectedSize - len(s)
	for i := 0; i < diff; i++ {
		s = fmt.Sprintf("%s%s", s, " ")
	}
	return s
}

func ContainsOneOfThese(containerString string, stringsToFind []string) bool {
	for _, stringToFind := range stringsToFind {
		if strings.Contains(containerString, stringToFind) {
			return true
		}
	}
	return false
}
