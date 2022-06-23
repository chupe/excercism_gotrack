package diffsquares

import "math"

func SquareOfSum(n int) int {
	var result float64
	for i := 0.0; i <= float64(n); i++ {
		result += i
	}
	return int(math.Pow(result, 2))
}

func SumOfSquares(n int) int {
	var result float64
	for i := 0.0; i <= float64(n); i++ {
		result += math.Pow(i, 2)
	}
	return int(result)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
