package montgomery

import (
	"math/big"
	"math/bits"
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

func multiply3(x, y, R, N *big.Int) *big.Int {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	// Convert x and y to Montgomery form: xR mod N, yR mod N
	xMont := redc2(x, rr, R, N)

	yMont := redc2(y, rr, R, N)

	// Perform Montgomery multiplication
	result := redc2(xMont, yMont, R, N)

	// Convert result out of Montgomery form: result * 1 mod N
	result = redc2(result, big.NewInt(1), R, N)

	return result
}

func multiply4(x, y, R, N *big.Int) *big.Int {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	// Convert x and y to Montgomery form: xR mod N, yR mod N
	xMont := redc3(x, rr, R, N)

	yMont := redc3(y, rr, R, N)

	// Perform Montgomery multiplication
	result := redc3(xMont, yMont, R, N)

	// Convert result out of Montgomery form: result * 1 mod N
	result = redc3(result, big.NewInt(1), R, N)

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

func redc2(x, y, R, N *big.Int) *big.Int {
	T := new(big.Int)
	yy := new(big.Int).Set(y)
	NI := newtonRaphsonInverse(N.Uint64())

	wordSize := 64

	s := R.BitLen() / wordSize

	for i := range s {
		var yi big.Word = 0
		if i < len(yy.Bits()) {
			yi = yy.Bits()[i]
		}
		t := new(big.Int).Mul(x, new(big.Int).SetUint64(uint64(yi)))
		T.Add(T, t)

		m := new(big.Int).Mul(T, new(big.Int).SetUint64(uint64(NI))).Uint64()

		T.Add(T, new(big.Int).Mul(new(big.Int).SetUint64(uint64(m)), N))
		T.Rsh(T, uint(wordSize))
	}
	if T.Cmp(N) >= 0 {
		T.Sub(T, N)
	}
	return T
}

func tobigInt(words []uint64) *big.Int {
	bits := make([]big.Word, len(words))
	for i, v := range words {
		bits[i] = big.Word(v)
	}
	result := new(big.Int)
	result.SetBits(bits)
	return result
}

func frombigInt(x *big.Int) []uint64 {
	words := x.Bits()
	result := make([]uint64, len(words))
	for i, w := range words {
		result[i] = uint64(w)
	}
	return result
}

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

// MontgomeryWords holds precomputed values for word-by-word Montgomery multiplication using redc3.
type MontgomeryWords struct {
	R  *big.Int  // R = 2^k
	N  *big.Int  // modulus (must be odd)
	RR *big.Int  // R² mod N (precomputed)
	NI uint64    // -N^(-1) mod 2^64 (precomputed via Newton-Raphson)
	S  int       // number of 64-bit words in R
	NN []uint64  // N as []uint64 (precomputed)
}

// NewMontgomeryWords creates a new MontgomeryWords instance with precomputed values.
func NewMontgomeryWords(R, N *big.Int) *MontgomeryWords {
	rr := new(big.Int).Mul(R, R)
	rr = rr.Mod(rr, N)

	wordSize := 64
	s := R.BitLen() / wordSize

	return &MontgomeryWords{
		R:  new(big.Int).Set(R),
		N:  new(big.Int).Set(N),
		RR: rr,
		NI: newtonRaphsonInverse(N.Uint64()),
		S:  s,
		NN: frombigInt(N),
	}
}

// Mul computes (x * y) mod N using word-by-word Montgomery multiplication.
func (m *MontgomeryWords) Mul(x, y *big.Int) *big.Int {
	// Convert to Montgomery form using precomputed R²
	xMont := m.redc(x, m.RR)
	yMont := m.redc(y, m.RR)

	// Montgomery multiplication
	result := m.redc(xMont, yMont)

	// Convert back from Montgomery form
	result = m.redc(result, big.NewInt(1))

	return result
}

// redc performs Montgomery reduction using word-by-word operations.
func (m *MontgomeryWords) redc(x, y *big.Int) *big.Int {
	T := make([]uint64, len(x.Bits())+len(y.Bits())+m.S+1)

	xx := frombigInt(x)
	yy := frombigInt(y)

	for i := range m.S {
		yi := uint64(0)
		if i < len(yy) {
			yi = yy[i]
		}

		// T += x * yi
		carry := uint64(0)
		for j, xj := range xx {
			hi, lo := bits.Mul64(xj, yi)
			s, c1 := bits.Add64(T[j], lo, 0)
			sum, c2 := bits.Add64(s, carry, 0)
			T[j] = sum
			carry = hi + c1 + c2
		}
		for k := len(xx); carry > 0 && k < len(T); k++ {
			sum, c := bits.Add64(T[k], carry, 0)
			T[k] = sum
			carry = c
		}

		// T += m * N
		mul := T[0] * m.NI
		carry = uint64(0)
		for j, nj := range m.NN {
			hi, lo := bits.Mul64(mul, nj)
			s, c1 := bits.Add64(T[j], lo, 0)
			sum, c2 := bits.Add64(s, carry, 0)
			T[j] = sum
			carry = hi + c1 + c2
		}
		for k := len(m.NN); carry > 0 && k < len(T); k++ {
			sum, c := bits.Add64(T[k], carry, 0)
			T[k] = sum
			carry = c
		}

		T = T[1:]
	}

	t := tobigInt(T)
	if t.Cmp(m.N) >= 0 {
		t.Sub(t, m.N)
	}
	return t
}

func redc3(x, y, R, N *big.Int) *big.Int {
	wordSize := 64
	s := R.BitLen() / wordSize

	// size is x*y+1, s is divided by word size
	T := make([]uint64, len(x.Bits())+len(y.Bits())+s+1)

	NI := newtonRaphsonInverse(N.Uint64())

	xx := frombigInt(x)
	yy := frombigInt(y)
	nn := frombigInt(N)

	for i := range s {
		yi := uint64(0)
		if i < len(yy) {
			yi = yy[i]
		}
		mulAddScalar(T, xx, yi)

		m := T[0] * NI
		mulAddScalar(T, nn, m)

		// Shift T right by wordSize bits
		T = T[1:]
	}

	t := tobigInt(T)
	if t.Cmp(N) >= 0 {
		t.Sub(t, N)
	}
	return t
}
