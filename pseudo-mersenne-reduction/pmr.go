package pseudomersennereduction

import "math"

const n = 8
const c = 5

var p = uint64(math.Exp2(n)) - c

func reduction(x uint64) uint64 {
	xx := x
	for xx >= 1<<n {
		high := xx >> n
		low := xx & (1<<n - 1)

		xx = high*c + low
	}
	for xx >= p {
		xx -= p
	}
	return xx
}
