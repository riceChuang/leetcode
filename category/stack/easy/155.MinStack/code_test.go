package leetcode

import "testing"

func TestStackFunc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StackFunc()
		})
	}
}
