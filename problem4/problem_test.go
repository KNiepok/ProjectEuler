package problem4

import (
	"testing"
)

func Test_largestPalindrome(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"largest palindrome for 2 numbers",
			args{
				2,
			},
			9009,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPalindrome(tt.args.digits); got != tt.want {
				t.Errorf("largestPalindrome() = %v, want %v", got, tt.want)
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
			"sample palindrome number #1",
			args{9009},
			true,
		},
		{
			"sample palindrome number #2",
			args{131},
			true,
		},
		{
			"sample palindrome number #3",
			args{13731},
			true,
		},
		{
			"sample not palindrome number #1",
			args{1234},
			false,
		},
		{
			"sample not palindrome number #2",
			args{13332},
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

func Test_Result(t *testing.T) {
	Result()
}
