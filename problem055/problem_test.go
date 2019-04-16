package problem055

import (
	"testing"
)

func Test_isLychrel(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 47",
			args{47},
			false,
		},
		{
			"test 349",
			args{349},
			false,
		},
		{
			"test 196",
			args{196},
			true,
		},
		{
			"test 4994",
			args{4994},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLychrel(tt.args.number); got != tt.want {
				t.Errorf("isLychrel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{121},
			true,
		},
		{
			"test",
			args{112},
			false,
		},
		{
			"test",
			args{1331},
			true,
		},
		{
			"test",
			args{1234},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.number); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{1331},
			1331,
		},
		{
			"test",
			args{1234},
			4321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInt(tt.args.n); got != tt.want {
				t.Errorf("reverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Result(t *testing.T) {
	Result()
}


