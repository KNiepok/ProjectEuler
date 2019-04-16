package problem010

import "testing"

func Test_sumOfPrimesBelow(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sum of primes below 10",
			args{
				10,
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPrimesBelow(tt.args.value); got != tt.want {
				t.Errorf("sumOfPrimesBelow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
