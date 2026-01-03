package rabin

import (
	"crypto/rand"
	"io"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func exp(p *big.Int) (s, d *big.Int) {
	d = new(big.Int).Sub(p, one)
	s = big.NewInt(0)

	tmp := new(big.Int)
	for {
		s.Add(s, one)
		d.Rsh(d, 1)
		if tmp.And(d, one).Cmp(one) == 0 {
			return
		}
	}
}

func IsPrime(n int64) (bool, error) {
	return isPrime(rand.Reader, big.NewInt(n), 20)
}

func isPrime(random io.Reader, p *big.Int, count int) (bool, error) {
	p = new(big.Int).Set(p)
	if p.Cmp(one) == 0 {
		return false, nil
	}
	if p.Cmp(two) == 0 {
		return true, nil
	}
	if new(big.Int).And(p, one).Cmp(zero) == 0 {
		// p%2==0(p&1==0)
		return false, nil
	}
	// p is odd number
	s, d := exp(p)
	max := new(big.Int).Sub(p, one)

	var (
		a      = new(big.Int)
		tmpa   = new(big.Int)
		tmpexp = new(big.Int)
	)

out:
	for i := 0; i < count; i++ {
		a.Set(zero)
		for a.Cmp(zero) == 0 {
			// a is not zero
			var err error
			a, err = rand.Int(random, p)
			if err != nil {
				return false, err
			}
		}
		tmpa = tmpa.Set(a).Exp(tmpa, d, p)
		if tmpa.Cmp(one) == 0 {
			// maybe prime if a^d==1(mod p)
			continue out
		}

		for r := big.NewInt(0); r.Cmp(s) < 0; r.Add(r, one) {
			tmpa.Set(a)
			tmpexp.Exp(two, r, nil).Mul(tmpexp, d)
			tmpa.Exp(tmpa, tmpexp, p)
			if tmpa.Cmp(max) == 0 {
				// maybe prime if a^(2^r*d)==-1==n-1(mod p)
				continue out
			}
		}
		// does not satisfy the prime number condition
		return false, nil
	}
	return true, nil
}
