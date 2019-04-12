package problem008

import "testing"

func Test_highestProduct(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"4 digits",
			args{4},
			5832,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestProduct(tt.args.digits); got != tt.want {
				t.Errorf("highestProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
