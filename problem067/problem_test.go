package problem067

import "testing"

func Test_maximumPath(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				triangle: [][]int{
					{3},
					{7,4},
					{2,4,6},
					{8,5,9,3},
				},
			},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPath(tt.args.triangle); got != tt.want {
				t.Errorf("maximumPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}

