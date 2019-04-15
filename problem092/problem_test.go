package problem092

import "testing"

func Test_nextItem(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{44},
			32,
		},
		{
			"test",
			args{32},
			13,
		},
		{
			"test",
			args{13},
			10,
		},
		{
			"test",
			args{85},
			89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextItem(tt.args.i); got != tt.want {
				t.Errorf("nextItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}

