package homework

import (
	"testing"
)

func TestJudgePalindrome(t *testing.T) {
	tests := [][]interface{}{
		{"中文回文", "杨恺是恺杨", true},
		{"中文非回文", "杨恺是杨恺", false},
		{"英文非回文", "are you ok?", false},
		{"英文回文", "MadamImAdam", true},
		{"空", "", true},
	}
	for _, line := range tests {
		t.Run(line[0].(string), func(t *testing.T) {
			s := line[1].(string)
			want := line[2].(bool)
			got := judgePalindrome(s)
			if got != want {
				t.Errorf("want %t, got %t\n", want, got)
			}
		})
	}
}
