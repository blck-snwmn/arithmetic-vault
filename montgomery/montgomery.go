// Package montgomery provides implementations of Montgomery multiplication,
// an efficient algorithm for modular arithmetic used in cryptography.
//
// Montgomery multiplication avoids costly division operations by transforming
// operands into "Montgomery form" (multiplied by R mod N), performing arithmetic
// in this domain, and then converting back.
//
// This package offers several implementations with different performance characteristics:
//   - MontgomeryBitwise: Basic bit-by-bit REDC algorithm
//   - MontgomeryCIOS: CIOS algorithm (word-by-word) using big.Int internally
//   - MontgomeryCIOSWords: CIOS algorithm using []uint64 for better performance
package montgomery

import (
	"math/big"
	"math/bits"
)

// MontgomeryBitwise holds precomputed values for bit-by-bit Montgomery multiplication.
type MontgomeryBitwise struct {
	R  *big.Int // R = 2^k
	N  *big.Int // modulus (must be odd)
	RR *big.Int // R² mod N (precomputed)
}

// NewMontgomeryBitwise creates a new MontgomeryBitwise instance with precomputed R² mod N.
func NewMontgomeryBitwise(R, N *big.Int) *MontgomeryBitwise {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)
	return &MontgomeryBitwise{
		R:  new(big.Int).Set(R),
		N:  new(big.Int).Set(N),
		RR: rr,
	}
}

// Mul computes (x * y) mod N using bit-by-bit Montgomery multiplication.
func (m *MontgomeryBitwise) Mul(x, y *big.Int) *big.Int {
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
func (m *MontgomeryBitwise) redc(x, y *big.Int) *big.Int {
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

// MontgomeryCIOS holds precomputed values for word-by-word Montgomery multiplication (CIOS algorithm).
type MontgomeryCIOS struct {
	R  *big.Int // R = 2^k
	N  *big.Int // modulus (must be odd)
	RR *big.Int // R² mod N (precomputed)
	NI uint64   // -N^(-1) mod 2^64 (precomputed via Newton-Raphson)
	S  int      // number of 64-bit words in R
}

// NewMontgomeryCIOS creates a new MontgomeryCIOS instance with precomputed values.
func NewMontgomeryCIOS(R, N *big.Int) *MontgomeryCIOS {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	wordSize := 64
	s := R.BitLen() / wordSize

	return &MontgomeryCIOS{
		R:  new(big.Int).Set(R),
		N:  new(big.Int).Set(N),
		RR: rr,
		NI: newtonRaphsonInverse(N.Uint64()),
		S:  s,
	}
}

// Mul computes (x * y) mod N using CIOS Montgomery multiplication.
func (m *MontgomeryCIOS) Mul(x, y *big.Int) *big.Int {
	// Convert to Montgomery form using precomputed R²
	xMont := m.redc(x, m.RR)
	yMont := m.redc(y, m.RR)

	// Montgomery multiplication
	result := m.redc(xMont, yMont)

	// Convert back from Montgomery form
	result = m.redc(result, big.NewInt(1))

	return result
}

// redc performs CIOS Montgomery reduction: (x * y * R⁻¹) mod N
func (m *MontgomeryCIOS) redc(x, y *big.Int) *big.Int {
	T := new(big.Int)
	yy := new(big.Int).Set(y)

	for i := 0; i < m.S; i++ {
		var yi big.Word = 0
		if i < len(yy.Bits()) {
			yi = yy.Bits()[i]
		}
		t := new(big.Int).Mul(x, new(big.Int).SetUint64(uint64(yi)))
		T.Add(T, t)

		mm := new(big.Int).Mul(T, new(big.Int).SetUint64(m.NI)).Uint64()

		T.Add(T, new(big.Int).Mul(new(big.Int).SetUint64(mm), m.N))
		T.Rsh(T, 64)
	}
	if T.Cmp(m.N) >= 0 {
		T.Sub(T, m.N)
	}
	return T
}

// MontgomeryCIOSWords holds precomputed values for CIOS Montgomery multiplication
// with optimized []uint64 representation for better performance.
type MontgomeryCIOSWords struct {
	R  *big.Int // R = 2^k
	N  *big.Int // modulus (must be odd)
	RR *big.Int // R² mod N (precomputed)
	NI uint64   // -N^(-1) mod 2^64 (precomputed via Newton-Raphson)
	S  int      // number of 64-bit words in R
	NN []uint64 // N as []uint64 (precomputed)
}

// NewMontgomeryCIOSWords creates a new MontgomeryCIOSWords instance with precomputed values.
func NewMontgomeryCIOSWords(R, N *big.Int) *MontgomeryCIOSWords {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	wordSize := 64
	s := R.BitLen() / wordSize

	return &MontgomeryCIOSWords{
		R:  new(big.Int).Set(R),
		N:  new(big.Int).Set(N),
		RR: rr,
		NI: newtonRaphsonInverse(N.Uint64()),
		S:  s,
		NN: frombigInt(N),
	}
}

// Mul computes (x * y) mod N using CIOS Montgomery multiplication.
func (m *MontgomeryCIOSWords) Mul(x, y *big.Int) *big.Int {
	// Convert to Montgomery form using precomputed R²
	xMont := m.redc(x, m.RR)
	yMont := m.redc(y, m.RR)

	// Montgomery multiplication
	result := m.redc(xMont, yMont)

	// Convert back from Montgomery form
	result = m.redc(result, big.NewInt(1))

	return result
}

// redc performs CIOS Montgomery reduction: (x * y * R⁻¹) mod N.
func (m *MontgomeryCIOSWords) redc(x, y *big.Int) *big.Int {
	T := make([]uint64, len(x.Bits())+len(y.Bits())+m.S+1)

	xx := frombigInt(x)
	yy := frombigInt(y)

	for i := range m.S {
		yi := uint64(0)
		if i < len(yy) {
			yi = yy[i]
		}

		mulAddScalar(T, xx, yi)

		// T += m * N
		mul := T[0] * m.NI
		mulAddScalar(T, m.NN, mul)

		T = T[1:]
	}

	t := tobigInt(T)
	if t.Cmp(m.N) >= 0 {
		t.Sub(t, m.N)
	}
	return t
}

// newtonRaphsonInverse computes -n^(-1) mod 2^64 using Newton-Raphson iteration.
//
// This value is used in Montgomery reduction to find the correction factor.
// The algorithm starts with x=1 (correct for 1 bit) and doubles precision
// each iteration via x = x * (2 - n*x), reaching 64-bit precision in 6 steps.
func newtonRaphsonInverse(n uint64) uint64 {
	x := uint64(1)

	x = x * (2 - n*x) // 2 bits
	x = x * (2 - n*x) // 4 bits
	x = x * (2 - n*x) // 8 bits
	x = x * (2 - n*x) // 16 bits
	x = x * (2 - n*x) // 32 bits
	x = x * (2 - n*x) // 64 bits
	return -x
}

// tobigInt converts a slice of uint64 words (little-endian) to *big.Int.
func tobigInt(words []uint64) *big.Int {
	bits := make([]big.Word, len(words))
	for i, v := range words {
		bits[i] = big.Word(v)
	}
	result := new(big.Int)
	result.SetBits(bits)
	return result
}

// frombigInt converts a *big.Int to a slice of uint64 words (little-endian).
func frombigInt(x *big.Int) []uint64 {
	words := x.Bits()
	result := make([]uint64, len(words))
	for i, w := range words {
		result[i] = uint64(w)
	}
	return result
}

// mulAddScalar computes T += arr * scalar using 64-bit word arithmetic.
//
// It performs a multiply-accumulate operation where each word of arr is
// multiplied by scalar, added to the corresponding word in T, with carry
// propagation handled correctly across word boundaries.
func mulAddScalar(T []uint64, arr []uint64, scalar uint64) {
	carry := uint64(0)
	for i, ai := range arr {
		hi, lo := bits.Mul64(ai, scalar)
		s, c1 := bits.Add64(T[i], lo, 0)
		sum, c2 := bits.Add64(s, carry, 0)
		T[i] = sum
		carry = hi + c1 + c2
	}
	for k := len(arr); carry > 0 && k < len(T); k++ {
		sum, c := bits.Add64(T[k], carry, 0)
		T[k] = sum
		carry = c
	}
}

// multiply computes (x * y) mod N using basic Montgomery multiplication.
//
// This is a straightforward implementation that converts operands to Montgomery
// form using direct multiplication (xR mod N), which requires expensive mod operations.
// For better performance, use Montgomery.Mul which uses REDC for the conversion.
func multiplyNaive(x, y, R, N *big.Int) *big.Int {
	// Convert x and y to Montgomery form: xR mod N, yR mod N
	xMont := new(big.Int).Mul(x, R)
	xMont = xMont.Mod(xMont, N)

	yMont := new(big.Int).Mul(y, R)
	yMont = yMont.Mod(yMont, N)

	// Perform Montgomery multiplication
	result := redcBitwise(xMont, yMont, R, N)

	// Convert result out of Montgomery form: result * 1 mod N
	result = redcBitwise(result, big.NewInt(1), R, N)

	return result
}

// redc performs bit-by-bit Montgomery reduction: (x * y * R⁻¹) mod N.
//
// The algorithm processes one bit at a time: if the LSB is 1, add N to make
// it even, then right-shift (divide by 2). After k iterations (where R = 2^k),
// the result is (x * y * R⁻¹) mod N.
func redcBitwise(x, y, R, N *big.Int) *big.Int {
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

