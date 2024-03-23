package longest_substring

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "Case 1",
			args: "abcabcbb",
			want: 3,
		},
		{
			name: "Case 2",
			args: "bbbbb",
			want: 1,
		},
		{
			name: "Case 3",
			args: "pwwkew",
			want: 3,
		},
		{
			name: "Case 4",
			args: "dvdf",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.args)

			if !cmp.Equal(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
