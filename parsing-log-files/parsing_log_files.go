package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[(ERR)])|^(\[(TRC)])|^(\[(DBG)])|^(\[(INF)])|^(\[(WRN)])|^(\[(FTL)])`)
	return re.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	var result int
	for _, line := range lines {
		result += len(re.FindAll([]byte(line), -1))
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return string(re.ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		if strings.Contains(line, "User ") {
			re := regexp.MustCompile(`User +\w*`)
			parsed := string(re.Find([]byte(line)))
			split := strings.Split(parsed, " ")
			user := split[len(split)-1]
			lines[i] = fmt.Sprintf("[USR] %s %s", user, line)
		}
	}
	return lines
}
