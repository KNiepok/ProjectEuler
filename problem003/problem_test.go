package problem003

import "testing"

func Test_largestPrimeFactor(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case #1",
			args{13195},
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPrimeFactor(tt.args.i); got != tt.want {
				t.Errorf("largestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_Result(t *testing.T) {
	Result()
}

