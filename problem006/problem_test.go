package problem006

import (
	"testing"
)

func Test_sumSquareDifference(t *testing.T) {
	type args struct {
		numbers int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"for first 10 numbers",
			args{10},
			2640,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSquareDifference(tt.args.numbers); got != tt.want {
				t.Errorf("sumSquareDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumOfSquares(t *testing.T) {
	type args struct {
		numbers int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"for first 10 numbers",
			args{10},
			385,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.args.numbers); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareSum(t *testing.T) {
	type args struct {
		numbers int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"for first 10 numbers",
			args{10},
			3025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareSum(tt.args.numbers); got != tt.want {
				t.Errorf("squareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
