package pseudomersennereduction

import (
	"math"
	"math/rand/v2"
	"testing"
)

func Test_reduction(t *testing.T) {
	tests := []struct {
		name string
		x    uint64
		want uint64
	}{
		// Below p: no reduction needed, input is returned as-is.
		{name: "x=0", x: 0, want: 0},
		{name: "x=1", x: 1, want: 1},
		{name: "x=p-1 (no reduction)", x: 250, want: 250},

		// Boundary band [p, 2^n): inner loop leaves xx unchanged because high=0.
		// The final `for xx >= p { xx -= p }` must bring these down.
		{name: "x=p", x: 251, want: 0},
		{name: "x=p+1", x: 252, want: 1},
		{name: "x=2^n-1", x: 255, want: 4},

		// Converges in a single inner-loop iteration.
		{name: "x=2^n (one iteration)", x: 256, want: 5},
		{name: "x=1000", x: 1000, want: 247},

		// Requires multiple inner-loop iterations:
		// 50000 = 195*256 + 80 -> 195*5+80 = 1055 -> 4*256+31 -> 4*5+31 = 51.
		{name: "x=50000 (multi-iteration)", x: 50000, want: 51},

		// Large inputs. Expected values derived from 2^n ≡ c (mod p).
		// 2^16 ≡ 5^2 = 25.
		{name: "x=2^16", x: 1 << 16, want: 25},
		// 2^32 ≡ 5^4 = 625 ≡ 123.
		{name: "x=2^32", x: 1 << 32, want: 123},
		// 2^63 ≡ 160 (from 2^64 ≡ 69, times the inverse of 2 mod p = 126).
		{name: "x=2^63", x: 1 << 63, want: 160},
		// 2^64 - 1 ≡ 68 (2^64 ≡ 69, minus 1).
		{name: "x=MaxUint64", x: math.MaxUint64, want: 68},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reduction(tt.x)
			if got != tt.want {
				t.Errorf("reduction() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Property test: reduction(x) must agree with the native `x % p` for random
// uint64 inputs. The seed is fixed to keep the test reproducible.
func Test_reduction_random(t *testing.T) {
	const iterations = 100000
	seed1, seed2 := uint64(0xA5A5A5A5A5A5A5A5), uint64(0x5A5A5A5A5A5A5A5A)
	r := rand.New(rand.NewPCG(seed1, seed2))

	for range iterations {
		x := r.Uint64()
		got := reduction(x)
		want := x % p
		if got != want {
			t.Fatalf("reduction(%d) = %d, want %d", x, got, want)
		}
	}
}
