package montgomery

import (
	"math/big"
	"testing"
	"testing/quick"
)

// testParams2048 returns 2048-bit test parameters for cryptographic-scale testing.
func testParams2048() (x, y, R, N *big.Int) {
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

func Test_newtonRaphsonInverse_maxUint64(t *testing.T) {
	t.Parallel()

	n := uint64(0xffffffffffffffff)
	ni := newtonRaphsonInverse(n)

	if ni != 0x0000000000000001 {
		t.Errorf("newtonRaphsonInverse(%#x) = %#x; want 0x1", n, ni)
	}
}

func Test_newtonRaphsonInverse_arbitraryOdd(t *testing.T) {
	t.Parallel()

	n := uint64(0xabcdef0123456789)
	ni := newtonRaphsonInverse(n)

	// n * ni should equal -1 (mod 2^64), i.e., 0xffffffffffffffff
	if n*ni != 0xffffffffffffffff {
		t.Errorf("newtonRaphsonInverse(%#x) = %#x; n*ni = %#x; want 0xffffffffffffffff", n, ni, n*ni)
	}
}

func Test_multiplyNaive(t *testing.T) {
	t.Parallel()

	x, y, R, N := testParams2048()
	result := multiplyNaive(x, y, R, N)
	expected := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

	if result.Cmp(expected) != 0 {
		t.Errorf("multiplyNaive: got %v; want %v", result, expected)
	}
}

func TestMontgomeryMul(t *testing.T) {
	t.Parallel()

	x2048, y2048, R2048, N2048 := testParams2048()
	N64, _ := new(big.Int).SetString("fffffffffffffffb", 16)
	R64 := new(big.Int).Lsh(big.NewInt(1), 64)

	tests := []struct {
		name string
		x    *big.Int
		y    *big.Int
		R    *big.Int
		N    *big.Int
	}{
		{
			name: "2048-bit cryptographic scale",
			x:    x2048,
			y:    y2048,
			R:    R2048,
			N:    N2048,
		},
		{
			name: "small values",
			x:    big.NewInt(7),
			y:    big.NewInt(11),
			R:    R64,
			N:    N64,
		},
		{
			name: "x equals zero",
			x:    big.NewInt(0),
			y:    big.NewInt(12345),
			R:    R64,
			N:    N64,
		},
		{
			name: "y equals zero",
			x:    big.NewInt(12345),
			y:    big.NewInt(0),
			R:    R64,
			N:    N64,
		},
		{
			name: "both zero",
			x:    big.NewInt(0),
			y:    big.NewInt(0),
			R:    R64,
			N:    N64,
		},
		{
			name: "x equals one",
			x:    big.NewInt(1),
			y:    big.NewInt(0x123456789abcdef),
			R:    R64,
			N:    N64,
		},
		{
			name: "y equals one",
			x:    big.NewInt(0x123456789abcdef),
			y:    big.NewInt(1),
			R:    R64,
			N:    N64,
		},
		{
			name: "x near N",
			x:    new(big.Int).Sub(N64, big.NewInt(1)),
			y:    big.NewInt(2),
			R:    R64,
			N:    N64,
		},
		{
			name: "y near N",
			x:    big.NewInt(2),
			y:    new(big.Int).Sub(N64, big.NewInt(1)),
			R:    R64,
			N:    N64,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			want := new(big.Int).Mod(new(big.Int).Mul(tc.x, tc.y), tc.N)

			t.Run("Bitwise", func(t *testing.T) {
				t.Parallel()
				m := NewMontgomeryBitwise(tc.R, tc.N)
				got := m.Mul(tc.x, tc.y)
				if got.Cmp(want) != 0 {
					t.Errorf("got %v, want %v", got, want)
				}
			})

			t.Run("CIOS", func(t *testing.T) {
				t.Parallel()
				m := NewMontgomeryCIOS(tc.R, tc.N)
				got := m.Mul(tc.x, tc.y)
				if got.Cmp(want) != 0 {
					t.Errorf("got %v, want %v", got, want)
				}
			})

			t.Run("CIOSWords", func(t *testing.T) {
				t.Parallel()
				m := NewMontgomeryCIOSWords(tc.R, tc.N)
				got := m.Mul(tc.x, tc.y)
				if got.Cmp(want) != 0 {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		})
	}
}

func TestMontgomeryMulProperty(t *testing.T) {
	t.Parallel()

	_, _, R, N := testParams2048()

	t.Run("Bitwise", func(t *testing.T) {
		t.Parallel()
		m := NewMontgomeryBitwise(R, N)

		err := quick.Check(func(xBytes, yBytes []byte) bool {
			x := new(big.Int).SetBytes(xBytes)
			y := new(big.Int).SetBytes(yBytes)
			x.Mod(x, N)
			y.Mod(y, N)

			got := m.Mul(x, y)
			want := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

			if got.Cmp(want) != 0 {
				return false
			}
			// result should be in range [0, N)
			return got.Sign() >= 0 && got.Cmp(N) < 0
		}, &quick.Config{MaxCount: 100})

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("CIOS", func(t *testing.T) {
		t.Parallel()
		m := NewMontgomeryCIOS(R, N)

		err := quick.Check(func(xBytes, yBytes []byte) bool {
			x := new(big.Int).SetBytes(xBytes)
			y := new(big.Int).SetBytes(yBytes)
			x.Mod(x, N)
			y.Mod(y, N)

			got := m.Mul(x, y)
			want := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

			if got.Cmp(want) != 0 {
				return false
			}
			return got.Sign() >= 0 && got.Cmp(N) < 0
		}, &quick.Config{MaxCount: 100})

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("CIOSWords", func(t *testing.T) {
		t.Parallel()
		m := NewMontgomeryCIOSWords(R, N)

		err := quick.Check(func(xBytes, yBytes []byte) bool {
			x := new(big.Int).SetBytes(xBytes)
			y := new(big.Int).SetBytes(yBytes)
			x.Mod(x, N)
			y.Mod(y, N)

			got := m.Mul(x, y)
			want := new(big.Int).Mod(new(big.Int).Mul(x, y), N)

			if got.Cmp(want) != 0 {
				return false
			}
			return got.Sign() >= 0 && got.Cmp(N) < 0
		}, &quick.Config{MaxCount: 100})

		if err != nil {
			t.Error(err)
		}
	})
}

func Benchmark_multiplyNaive(b *testing.B) {
	x, y, R, N := testParams2048()

	for b.Loop() {
		multiplyNaive(x, y, R, N)
	}
}

func BenchmarkMontgomeryMul(b *testing.B) {
	x, y, R, N := testParams2048()

	b.Run("Bitwise", func(b *testing.B) {
		m := NewMontgomeryBitwise(R, N)
		for b.Loop() {
			m.Mul(x, y)
		}
	})

	b.Run("CIOS", func(b *testing.B) {
		m := NewMontgomeryCIOS(R, N)
		for b.Loop() {
			m.Mul(x, y)
		}
	})

	b.Run("CIOSWords", func(b *testing.B) {
		m := NewMontgomeryCIOSWords(R, N)
		for b.Loop() {
			m.Mul(x, y)
		}
	})
}
