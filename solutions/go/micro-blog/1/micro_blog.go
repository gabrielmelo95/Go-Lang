package microblog

import (
	"strings"
	"unicode/utf8"
)

func Truncate(phrase string) string {
	var truncate strings.Builder
	cnt := 0
	for len(phrase) > 0 && cnt < 5 {
		r, size := utf8.DecodeRuneInString(phrase)
		phrase = phrase[size:]
		truncate.WriteString(string(r))
		cnt++
	}
	return truncate.String()
}
