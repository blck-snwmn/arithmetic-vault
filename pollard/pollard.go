package pollard

import "math/big"

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

func gcd(l, r *big.Int) *big.Int {
	for r.Cmp(zero) != 0 {
		l, r = r, new(big.Int).Mod(l, r)
	}
	return l
}

func floydo(n *big.Int) *big.Int {
	return inner_floydo(n, big.NewInt(1))
}

func inner_floydo(n, c *big.Int) *big.Int {
	f := func(x, n *big.Int) *big.Int {
		sq := new(big.Int).Mul(x, x)
		ad := sq.Add(sq, c)
		return ad.Mod(ad, n)
	}

	subAbs := func(l, r *big.Int) *big.Int {
		if l.Cmp(r) > 0 {
			return new(big.Int).Sub(l, r)
		}
		return new(big.Int).Sub(r, l)
	}

	x := big.NewInt(1)
	y := big.NewInt(1)
	d := big.NewInt(1)
	for d.Cmp(one) == 0 {
		x = f(x, n)
		y = f(f(y, n), n)
		d = gcd(subAbs(x, y), n)
	}
	return d
}
