package problem001

import "testing"

func Test_sumOfNumbers(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sum of 3,5,6 and 9 is equal to 23",
			args{10},
			23,
		},
		{
			"sum of 3,5,6,9,10,12,15,18 is equal to 78",
			args{20},
			78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfNumbers(tt.args.limit); got != tt.want {
				t.Errorf("sumOfNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
