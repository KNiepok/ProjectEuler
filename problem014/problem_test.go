package problem014

import "testing"

func Test_collatzChainLength(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"chain for 13: 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1 has 10 items",
			args{13},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMapped()
			if got := collatzChainLength(tt.args.i); got != tt.want {
				t.Errorf("collatzChainLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
