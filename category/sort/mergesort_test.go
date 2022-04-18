package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		question []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{
				[]int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.question); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
