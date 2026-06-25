package isogram

import "strings"

func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")
	head := word[0]
	tail := word[1:]
	repeat := 0
	for j := 0; j < len(word)-1; j++ {
		for i := 0; i < len(tail); i++ {
			if head == tail[i] {
				repeat++
				if repeat > 0 {
					return false
				}
			}
		}
		head = tail[0]
		tail = tail[1:]
	}
	return true
}
