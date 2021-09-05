package palindromenumber9

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{
				x: 123,
			},
			false,
		},
		{
			"test",
			args{
				x: 1234321,
			},
			true,
		},
		{
			"test",
			args{
				x: -1234321,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
