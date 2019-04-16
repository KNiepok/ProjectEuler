package problem081

import "testing"

func Test_shortestPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"5x5 grid",
			args{
				[][]int{
					{131, 673, 234, 103, 18},
					{201, 96, 342, 965, 150},
					{630, 803, 746, 422, 111},
					{537, 699, 497, 121, 956},
					{805, 732, 524, 37, 331},
				},
			},
			2427,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.grid); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}
