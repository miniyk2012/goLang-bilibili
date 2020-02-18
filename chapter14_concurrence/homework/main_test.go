package main

import (
	"testing"
	"fmt"
)

func TestNumSum(t *testing.T) {
	testCases := [][]int {
		{345, 13},
		{1, 1},
		{102, 3},
	}
	for _, testCase := range testCases {
		num, total := testCase[0], testCase[1]
		t.Run(fmt.Sprintf("%d num sum is %d", num, total), func(t *testing.T)() {
			result := calNumSum(int64(346))
			if result != 13 {
				t.Errorf("result=%d, expected=%d", result, 13)
			}
		})
	}
}
