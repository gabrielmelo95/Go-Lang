package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	word := ""
	subject = strings.ToLower(subject)
	for i := range candidates {
		word = strings.ToLower(candidates[i])
		if len(subject) != len(candidates[i]) || word == subject {
			continue
		}
		for _, r := range subject {
			if unicode.IsUpper(r) {
				unicode.ToLower(r)
			}
			if strings.ContainsRune(strings.ToLower(candidates[i]), r) {
				word = strings.Replace(word, string(r), "", 1)
			}
		}
		if word == "" {
			anagrams = append(anagrams, candidates[i])
		}
	}
	return anagrams
}
