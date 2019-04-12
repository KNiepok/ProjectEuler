package problem031

import "testing"

func Test_differentWays(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countWays()
		})
	}
}
