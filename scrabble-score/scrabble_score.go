package scrabble

import "strings"

// var (
// 	onePoint = [...]rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
//
// )
const (
	onePoint   = "AEIOULNRST"
	twoPoint   = "DG"
	threePoint = "BCMP"
	fourPoint  = "FHVWY"
	fivePoint  = "K"
	eightPoint = "JX"
	tenPoint   = "QZ"
)

/*Score returns integer value, based on argument string. Each letter in the argument has particular value.
Returned value is sum of these values.*/
func Score(input string) int {
	if input == "" {
		return 0
	}
	result := 0
	input = strings.ToUpper(input)
	for _, letter := range input {
		result += getValue(letter)
	}
	return result
}

// getValue returns integer value of the argument.
func getValue(letter rune) int {
	switch {
	case strings.ContainsRune(onePoint, letter):
		return 1
	case strings.ContainsRune(twoPoint, letter):
		return 2
	case strings.ContainsRune(threePoint, letter):
		return 3
	case strings.ContainsRune(fourPoint, letter):
		return 4
	case strings.ContainsRune(fivePoint, letter):
		return 5
	case strings.ContainsRune(eightPoint, letter):
		return 8
	case strings.ContainsRune(tenPoint, letter):
		return 10
	default:
		return 0
	}
}
