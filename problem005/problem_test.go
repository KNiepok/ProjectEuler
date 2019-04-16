package problem005

import (
	"reflect"
	"testing"
)

func Test_smallestNumberDivisible(t *testing.T) {
	type args struct {
		numbers int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"numbers from 1 to 10",
			args{10},
			2520,
		},
		{
			"numbers from 1 to 3",
			args{3},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumberDivisible(tt.args.numbers); got != tt.want {
				t.Errorf("smallestNumberDivisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrimes(t *testing.T) {
	type args struct {
		upperLimit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{10},
			[]int{2, 3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrimes(tt.args.upperLimit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
