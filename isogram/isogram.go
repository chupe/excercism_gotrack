package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, r := range word {
		reg := regexp.MustCompile("[A-z]")
		if reg.MatchString(string(r)) &&
			strings.Count(word, strings.ToLower(string(r))) > 1 {
			return false
		}
	}
	return true
}
