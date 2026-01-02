package montgomery

import (
	"math/big"
)

func multiply(x, y, R, N *big.Int) *big.Int {
	// Convert x and y to Montgomery form: xR mod N, yR mod N
	xMont := new(big.Int).Mul(x, R)
	xMont = xMont.Mod(xMont, N)

	yMont := new(big.Int).Mul(y, R)
	yMont = yMont.Mod(yMont, N)

	// Perform Montgomery multiplication
	result := redc(xMont, yMont, R, N)

	// Convert result out of Montgomery form: result * 1 mod N
	result = redc(result, big.NewInt(1), R, N)

	return result
}

func multiply2(x, y, R, N *big.Int) *big.Int {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	// Convert x and y to Montgomery form: xR mod N, yR mod N
	xMont := redc(x, rr, R, N)

	yMont := redc(y, rr, R, N)

	// Perform Montgomery multiplication
	result := redc(xMont, yMont, R, N)

	// Convert result out of Montgomery form: result * 1 mod N
	result = redc(result, big.NewInt(1), R, N)

	return result
}

func redc(x, y, R, N *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)

	// Loop k times for Montgomery reduction where R = 2^k
	for r := 1; r < R.BitLen(); r++ {
		if result.Bit(0) == 1 {
			result.Add(result, N)
		}
		result.Rsh(result, 1)
	}
	if result.Cmp(N) >= 0 {
		result.Sub(result, N)
	}

	return result
}
