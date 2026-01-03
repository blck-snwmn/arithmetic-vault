package montgomery

import (
	"math/big"
	"testing"
)

func Test_newtonRaphsonInverse(t *testing.T) {
	{
		n := uint64(0xffffffffffffffff)
		ni := newtonRaphsonInverse(n)

		expected := uint64(0x0000000000000001)
		if ni != expected {
			t.Errorf("newtonRaphsonInverse(%v) = %v; want %v", n, ni, expected)
		}
	}
	{
		n := uint64(0xabcdef0123456789)
		ni := newtonRaphsonInverse(n)

		if n*ni != 0xffffffffffffffff {
			t.Errorf("newtonRaphsonInverse(%v) = %v; n*ni = %v; want %v", n, ni, n*ni, "0xffffffffffffffff")
		}
	}
}

func Test_multiplyNaive(t *testing.T) {
	x, y, R, N := testParams()

	result := multiplyNaive(x, y, R, N)

	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("redc(%v, %v, %v, %v) = %v; want %v", x, y, R, N, result, expected)
	}
}

func testParams() (x, y, R, N *big.Int) {
	x, _ = new(big.Int).SetString(""+
		"a3b2c1d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2"+
		"c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4"+
		"e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6"+
		"a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8"+
		"c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0"+
		"e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2"+
		"a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4"+
		"c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6", 16)

	y, _ = new(big.Int).SetString(""+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321", 16)

	N, _ = new(big.Int).SetString(""+
		"d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5"+
		"f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7"+
		"b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9"+
		"d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1"+
		"f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3"+
		"b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5"+
		"d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7"+
		"f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9", 16)

	R = new(big.Int).Lsh(big.NewInt(1), 2048)
	return
}

func Benchmark_multiplyNaive(b *testing.B) {
	x, y, R, N := testParams()

	for b.Loop() {
		multiplyNaive(x, y, R, N)
	}
}

func TestMontgomeryBitwise_Mul(t *testing.T) {
	x, y, R, N := testParams()

	m := NewMontgomeryBitwise(R, N)
	result := m.Mul(x, y)

	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("MontgomeryBitwise.Mul(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

func Benchmark_MontgomeryBitwise_Mul(b *testing.B) {
	x, y, R, N := testParams()
	m := NewMontgomeryBitwise(R, N)

	for b.Loop() {
		m.Mul(x, y)
	}
}

func TestMontgomeryCIOS_Mul(t *testing.T) {
	x, y, R, N := testParams()

	m := NewMontgomeryCIOS(R, N)
	result := m.Mul(x, y)

	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("MontgomeryCIOS.Mul(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

func Benchmark_MontgomeryCIOS_Mul(b *testing.B) {
	x, y, R, N := testParams()
	m := NewMontgomeryCIOS(R, N)

	for b.Loop() {
		m.Mul(x, y)
	}
}

func TestMontgomeryCIOSWords_Mul(t *testing.T) {
	x, y, R, N := testParams()

	m := NewMontgomeryCIOSWords(R, N)
	result := m.Mul(x, y)

	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("MontgomeryCIOSWords.Mul(%v, %v) = %v; want %v", x, y, result, expected)
	}
}

func Benchmark_MontgomeryCIOSWords_Mul(b *testing.B) {
	x, y, R, N := testParams()
	m := NewMontgomeryCIOSWords(R, N)

	for b.Loop() {
		m.Mul(x, y)
	}
}
