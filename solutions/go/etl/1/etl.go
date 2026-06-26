package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for v, s := range in {
		for _, c := range s {
			c = strings.ToLower(c)
			res[c] = v
		}
	}
	return res
}
