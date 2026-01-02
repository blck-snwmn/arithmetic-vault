package montgomery

import (
	"math/big"
	"testing"
)

func Test_multiply(t *testing.T) {
	x, _ := new(big.Int).SetString(""+
		"a3b2c1d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2"+
		"c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4"+
		"e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6"+
		"a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8"+
		"c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0"+
		"e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2"+
		"a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4"+
		"c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6", 16)

	y, _ := new(big.Int).SetString(""+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321", 16)

	// N must be odd for Montgomery reduction
	N, _ := new(big.Int).SetString(""+
		"d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5"+
		"f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7"+
		"b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9"+
		"d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1"+
		"f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3"+
		"b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5"+
		"d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7"+
		"f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9", 16)

	R := new(big.Int).Lsh(big.NewInt(1), 2048) // R = 2^2048

	result := multiply(x, y, R, N)

	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("redc(%v, %v, %v, %v) = %v; want %v", x, y, R, N, result, expected)
	}
}

func Test_multiply2(t *testing.T) {
	x, _ := new(big.Int).SetString(""+
		"a3b2c1d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2"+
		"c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4"+
		"e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6"+
		"a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8"+
		"c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0"+
		"e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2"+
		"a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4"+
		"c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6", 16)

	y, _ := new(big.Int).SetString(""+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"+
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"+
		"fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321", 16)

	// N must be odd for Montgomery reduction
	N, _ := new(big.Int).SetString(""+
		"d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5"+
		"f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7"+
		"b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9"+
		"d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1"+
		"f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3"+
		"b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5"+
		"d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7"+
		"f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9", 16)

	R := new(big.Int).Lsh(big.NewInt(1), 2048) // R = 2^2048

	result := multiply2(x, y, R, N)

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

func Benchmark_multiply(b *testing.B) {
	x, y, R, N := testParams()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		multiply(x, y, R, N)
	}
}

func Benchmark_multiply2(b *testing.B) {
	x, y, R, N := testParams()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		multiply2(x, y, R, N)
	}
}
