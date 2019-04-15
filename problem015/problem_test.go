package problem015

import "testing"

func Test_letticePaths(t *testing.T) {
	type args struct {
		x, y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"2x2 grid should have 6 paths",
			args{2, 2},
			6,
		},
		{
			"2x1 grid should have 3 paths",
			args{2, 1},
			3,
		},
		{
			"3x0 grid should have 1 path",
			args{3, 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initMap()
			if got := letticePaths(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("leticePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
