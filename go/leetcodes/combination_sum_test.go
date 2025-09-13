package leetcodesgo

import (
	"reflect"
	"strings"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	stringAnswers := strings.Join([]string{
		"[1,1,1,1,1,1,1]",
		"[1,1,1,1,1,2]",
		"[1,1,1,1,3]",
		"[1,1,1,2,2]",
		"[1,1,2,3]",
		"[1,2,2,2]",
		"[1,3,3]",
		"[2,2,3]",
	}, ",")

	expectedResult, err := parseIntegerArrayArray(stringAnswers)
	if err != nil {
		t.Errorf("Parsing error! %s\n", err)
	}

	result := combinationSum([]int{1, 2, 3}, 7)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
