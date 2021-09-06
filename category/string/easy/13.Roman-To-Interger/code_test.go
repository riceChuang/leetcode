package string

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				s: "III",
			},
			3,
		},
		{
			"test",
			args{
				s: "IV",
			},
			4,
		},
		{
			"test",
			args{
				s: "LVIII",
			},
			58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
