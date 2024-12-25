package e3

import "strings"

func Prefixer(prefix string) func(string) string {
	return func(s string) string {
		elems := []string{prefix, s}
		return strings.Join(elems, " ")
	}

}
