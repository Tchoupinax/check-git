package utils

import "fmt"

func AddSpaceToEnd(s string, size int) string {
	var diff = size - len(s)

	for i := 0; i < diff; i++ {
		s = fmt.Sprintf("%s%s", s, " ")
	}

	return s
}
