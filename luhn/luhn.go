package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.Replace(id, " ", "", -1)
	if len(id) < 2 {
		return false
	}

	var total int
	length := len(id)
	for i, r := range id {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if (length+i)%2 == 0 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
	}

	return total%10 == 0
}
