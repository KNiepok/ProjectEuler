package problem002

import "testing"

func Test_evenFibonacciSum(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case",
			args{
				90,
			},
			44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenFibonacciSum(tt.args.limit); got != tt.want {
				t.Errorf("evenFibonacciSum() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_Result(t *testing.T) {
	Result()
}
