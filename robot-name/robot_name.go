package robotname

import (
	"errors"
	"fmt"
)

type Robot struct {
	name string
}

func generateName() func() (string, error) {
	c1Counter := int32(0)
	c2Counter := int32(0)
	digits := 0

	return func() (string, error) {
		if digits == 1000 {
			digits = 0
			c2Counter++
		}
		if c2Counter > 25 {
			c1Counter++
			c2Counter = 0
		}
		if c1Counter > 25 {
			return "", errors.New("namespace is exhausted")
		}

		c1 := 'A' + c1Counter
		c2 := 'A' + c2Counter
		result := fmt.Sprintf("%s%s%03d", string(c1), string(c2), digits)

		digits++
		return result, nil
	}
}

var getName = generateName()

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := getName()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
