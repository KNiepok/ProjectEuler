package problem5

import "testing"

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
