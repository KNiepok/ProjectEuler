package problem025

import "testing"

func Test_nDigitFibonacciNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first 3 digit fib",
			args{3},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nDigitFibonacciNumber(tt.args.n); got != tt.want {
				t.Errorf("nDigitFibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}

