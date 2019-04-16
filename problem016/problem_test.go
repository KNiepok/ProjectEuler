package problem016

import "testing"

func Test_powerDigitSum(t *testing.T) {
	type args struct {
		power int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2^15 = 32768 => 3 + 2 + 7 + 6 + 8 = 26.",
			args{15},
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerDigitSum(tt.args.power); got != tt.want {
				t.Errorf("powerDigitSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
