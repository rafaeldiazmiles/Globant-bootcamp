package calculor

import "errors"

func Add(a, b float64) float64 {
	return a + b
}

func Subs(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("dividing by 0?? Have you lost your mind???")
	}

	return a / b, nil

}
