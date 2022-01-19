package factorial_test

import (
	"testing"

	. "go-bits/variable-names/factorial"
)

func Test_PositiveFactorial(t *testing.T) {
	type args struct {
		number uint8
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"0", args{0}, 1},
		{"1", args{1}, 1},
		{"2", args{2}, 2},
		{"3", args{3}, 6},
		{"4", args{4}, 24},
		{"5", args{5}, 120},
		{"6", args{6}, 720},
		{"7", args{7}, 5040},
		{"55", args{55}, 6711489344688881664},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositiveFactorial(tt.args.number); got != tt.want {
				t.Errorf("PositiveFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
