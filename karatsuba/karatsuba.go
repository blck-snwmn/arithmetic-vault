package karatsuba

import (
	"math"
)

func len10(x uint) uint {
	digits := 0
	for n := x; n > 0; n /= 10 {
		digits++
	}
	return uint(digits)
}
func mul(x, y uint) uint {
	lx := len10(x) / 2
	ly := len10(y) / 2

	l := max(lx, ly)

	tt := math.Pow10(int(l))
	t := uint(tt)

	a := x / t
	b := x % t
	c := y / t
	d := y % t

	ac := a * c
	bd := b * d

	abcd := (a + b) * (c + d)

	return ac*t*t + (abcd-ac-bd)*t + bd
}

func mulRec(x, y uint) uint {
	if x < 1000 || y < 1000 {
		return x * y
	}
	math.Log10(float64(x))
	lx := len10(x) / 2
	ly := len10(y) / 2

	l := max(lx, ly)

	tt := math.Pow10(int(l))
	t := uint(tt)

	a := x / t
	b := x % t
	c := y / t
	d := y % t

	ac := mulRec(a, c)
	bd := mulRec(b, d)

	abcd := mulRec((a + b), (c + d))

	return ac*t*t + (abcd-ac-bd)*t + bd
}
