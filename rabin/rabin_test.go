package rabin

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func Test_isPrime(t *testing.T) {
	type args struct {
		p *big.Int
	}
	tests := []struct {
		args    args
		want    bool
		wantErr bool
	}{
		{args: args{p: big.NewInt(1)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(2)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(3)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(11)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(41)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(100)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(111)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(127)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(1150)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(1151)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(1153)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(1155)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(1223)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(2295)}, want: false, wantErr: false},
		{args: args{p: big.NewInt(4423)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(6427)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(7603)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(8191)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(18041)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(24239)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(40507)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(68963)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(131071)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(524287)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(6700417)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(2147483647)}, want: true, wantErr: false},
		{args: args{p: big.NewInt(67280421310721)}, want: true, wantErr: false},
		{args: args{p: new(big.Int).Sub(new(big.Int).Exp(two, big.NewInt(127), nil), one)}, want: true, wantErr: false},
		{args: args{p: new(big.Int).Exp(two, big.NewInt(4423), nil)}, want: false, wantErr: false},
		{args: args{p: new(big.Int).Sub(new(big.Int).Exp(two, big.NewInt(4423), nil), one)}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("input=%v", tt.args.p), func(t *testing.T) {
			t.Parallel()
			got, err := isPrime(rand.Reader, tt.args.p, 20)
			if (err != nil) != tt.wantErr {
				t.Errorf("isPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exp(t *testing.T) {
	type args struct {
		p *big.Int
	}
	tests := []struct {
		name  string
		args  args
		wantS *big.Int
		wantD *big.Int
	}{
		{
			name:  "11-1=2^1*5",
			args:  args{p: big.NewInt(11)},
			wantS: big.NewInt(1),
			wantD: big.NewInt(5),
		},
		{
			name:  "31-1=2^1*15",
			args:  args{p: big.NewInt(31)},
			wantS: big.NewInt(1),
			wantD: big.NewInt(15),
		},
		{
			name:  "25-1=2^3*3",
			args:  args{p: big.NewInt(25)},
			wantS: big.NewInt(3),
			wantD: big.NewInt(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotD := exp(tt.args.p)
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("exp() gotS = %v, want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotD, tt.wantD) {
				t.Errorf("exp() gotD = %v, want %v", gotD, tt.wantD)
			}
		})
	}
}

func TestXxx(t *testing.T) {
	// d := new(big.Int).Sub(big.NewInt(11), one)
	// s := big.NewInt(0)
	// fmt.Println("a", d, s)
	// s.Add(s, one)
	// fmt.Println("b", d, s)
	// d.Rsh(d, 1)
	// fmt.Println("c", d, s)
	// x := new(big.Int).And(d, one)
	// fmt.Println("d", d, s, x)
	// fmt.Println("x", d)
	two := big.NewInt(2)
	tree := big.NewInt(3)
	x := new(big.Int).Mod(two, tree)
	fmt.Println(x)
}
