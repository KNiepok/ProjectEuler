package problem007

import (
	"testing"
)

func Test_nthPrime(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{6},
			13,
		},
		{
			"test",
			args{7},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPrime(tt.args.i); got != tt.want {
				t.Errorf("nthPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
