package leetcode

import "testing"

func Test_factorial_DP(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factorial_DP()
		})
	}
}

func Test_fibonacci_loop(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				10,
			},
			55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci_loop(tt.args.target); got != tt.want {
				t.Errorf("fibonacci_loop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci_DP(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				10,
			},
			55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci_DP(tt.args.target); got != tt.want {
				t.Errorf("fibonacci_DP() = %v, want %v", got, tt.want)
			}
		})
	}
}
