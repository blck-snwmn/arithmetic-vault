package karatsuba

import "testing"

func Test_len(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    uint
		want uint
	}{
		{"2桁", 10, 2},
		{"3桁", 100, 3},
		{"6桁", 412412, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len10(tt.x)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    uint
		y    uint
		want uint
	}{
		{"3桁x3桁", 101, 202, 20402},
		{"5桁x3桁", 10101, 202, 2040402},
		{"5桁x5桁", 12345, 54321, 670592745},
		{"6桁x6桁", 123456, 654321, 80779853376},
		{"7桁x7桁", 9999999, 9999999, 99999980000001},
		{"8桁x3桁", 12345678, 999, 12333332322},
		{"桁数が大きく異なる", 99999999, 2, 199999998},
		{"片方が1桁", 7, 8765432, 61358024},
		{"10桁x10桁", 1234567890, 9876543210, 12193263111263526900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mul(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mulRec(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    uint
		y    uint
		want uint
	}{
		{"3桁x3桁", 101, 202, 20402},
		{"5桁x3桁", 10101, 202, 2040402},
		{"5桁x5桁", 12345, 54321, 670592745},
		{"6桁x6桁", 123456, 654321, 80779853376},
		{"7桁x7桁", 9999999, 9999999, 99999980000001},
		{"8桁x3桁", 12345678, 999, 12333332322},
		{"桁数が大きく異なる", 99999999, 2, 199999998},
		{"片方が1桁", 7, 8765432, 61358024},
		{"10桁x10桁", 1234567890, 9876543210, 12193263111263526900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mulRec(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
