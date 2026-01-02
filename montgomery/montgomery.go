package montgomery

import (
	"math/big"
)

// Montgomery holds precomputed values for Montgomery multiplication.
type Montgomery struct {
	R  *big.Int // R = 2^k
	N  *big.Int // modulus (must be odd)
	RR *big.Int // R² mod N (precomputed)
}

// NewMontgomery creates a new Montgomery instance with precomputed R² mod N.
func NewMontgomery(R, N *big.Int) *Montgomery {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)
	return &Montgomery{
		R:  new(big.Int).Set(R),
		N:  new(big.Int).Set(N),
		RR: rr,
	}
}

// Mul computes (x * y) mod N using Montgomery multiplication.
func (m *Montgomery) Mul(x, y *big.Int) *big.Int {
	// Convert to Montgomery form using precomputed R²
	xMont := m.redc(x, m.RR)
	yMont := m.redc(y, m.RR)

	// Montgomery multiplication
	result := m.redc(xMont, yMont)

	// Convert back from Montgomery form
	result = m.redc(result, big.NewInt(1))

	return result
}

// redc performs Montgomery reduction: (x * y * R⁻¹) mod N
func (m *Montgomery) redc(x, y *big.Int) *big.Int {
	result := new(big.Int).Mul(x, y)

	// Loop k times for Montgomery reduction where R = 2^k
	for r := 1; r < m.R.BitLen(); r++ {
		if result.Bit(0) == 1 {
			result.Add(result, m.N)
		}
		result.Rsh(result, 1)
	}
	if result.Cmp(m.N) >= 0 {
		result.Sub(result, m.N)
	}

	return result
}

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
