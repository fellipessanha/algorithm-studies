package leetcodesgo

import (
	"testing"
)

func testMaxSumWithOneSubarrayDeletion(t *testing.T) {
	instances := map[string]int{
		"1, -2, 0, 3":  4,
		"1, -2, -2, 3": 3,
		"-1, -1, -1":   -1,
	}
	for input, answer := range instances {
		values, _ := parseIntegerArray(input)
		if maximumSum(values) != answer {
			t.Errorf("%v got wrong answer!", input)
		}
	}
}
