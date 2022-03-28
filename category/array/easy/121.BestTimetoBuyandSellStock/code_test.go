package leetcode

import "testing"

func Test_maxProfit_new(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_new(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_new() = %v, want %v", got, tt.want)
			}
		})
	}
}
