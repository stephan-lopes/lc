package twosums

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwoSums(t *testing.T) {

	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args input
		want []int
	}{
		{
			name: "Case 1",
			args: input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Case 2",
			args: input{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Case 3",
			args: input{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.args.nums, tt.args.target)

			if !cmp.Equal(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
