package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	matchLogExp := `\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\]`
	sl := regexp.MustCompile(matchLogExp).FindStringIndex(text)
	if sl != nil {
		return sl[0] == 0
	}
	return false
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`\<[~*=-]*\>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)\".*password.*\"`)
	count := 0
	for _, line := range lines {
		if re.FindString(line) != "" {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line[0-9]*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	re := regexp.MustCompile(`User\s+(\S+)`)

	var result []string

	for _, line := range lines {
		match := re.FindStringSubmatch(line)

		if match != nil {
			username := match[1]
			line = fmt.Sprintf("[USR] %s %s", username, line)
		}

		result = append(result, line)
	}

	return result
}
