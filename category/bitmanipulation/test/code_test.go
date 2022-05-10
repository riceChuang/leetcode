package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		cars []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{
				[]string{"100", "110", "010", "011", "100"},
			},
			[]int{2, 3, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.cars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
