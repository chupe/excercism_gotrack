package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number out of range")
	}
	//return uint64(math.Pow(2, float64(number-1))), nil
	return uint64(1) << (number - 1), nil
}

func Total() uint64 {
	//var result uint64
	//for i := 1; i <= 64; i++ {
	//	square, _ := Square(i)
	//	//if err != nil {
	//	//	break
	//	//}
	//	result += square
	//}
	//return result

	const maxUint = 1<<64 - 1
	return maxUint
}
