package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	q := []int{12, 85, 25, 16, 34, 23, 49, 95, 17, 61}
	type args struct {
		question []int
		begin    int
		end      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test",
			args{
				question: q,
				begin:    0,
				end:      len(q) - 1,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.question, tt.args.begin, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
