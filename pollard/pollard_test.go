package pollard

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_gcd(t *testing.T) {
	type args struct {
		l *big.Int
		r *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "gcd(10, 5) = 5",
			args: args{
				l: big.NewInt(10),
				r: big.NewInt(5),
			},
			want: big.NewInt(5),
		},
		{
			name: "gcd(10, 4) = 2",
			args: args{
				l: big.NewInt(10),
				r: big.NewInt(4),
			},
			want: big.NewInt(2),
		},
		{
			name: "gcd(10, 3) = 1",
			args: args{
				l: big.NewInt(10),
				r: big.NewInt(3),
			},
			want: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floydo(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "floydo(15) = 3",
			args: args{
				n: big.NewInt(15),
			},
			want: big.NewInt(3),
		},
		{
			name: "floydo(221) = 13",
			args: args{
				n: big.NewInt(221),
			},
			want: big.NewInt(13),
		},
		{
			name: "floydo(8051) = 97",
			args: args{
				n: big.NewInt(8051),
			},
			want: big.NewInt(97),
		},
		{
			name: "floydo(1223) = 1223",
			args: args{
				n: big.NewInt(1223),
			},
			want: big.NewInt(1223),
		},
		{
			name: "floydo(29) = 29",
			args: args{
				n: big.NewInt(29),
			},
			want: big.NewInt(29),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floydo(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floydo() = %v, want %v", got, tt.want)
			}
		})
	}
}
