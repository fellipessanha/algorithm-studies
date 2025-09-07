package leetcodesgo

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func parseStringAnswer(entry []string) ([][]int, error) {
	parsedEntry := make([][]int, 0, len(entry))
	for _, entryLine := range entry {
		answerParcel, err := parseStringAnswerLine(entryLine)
		if err != nil {
			return [][]int{}, err
		}
		parsedEntry = append(parsedEntry, answerParcel)
	}
	return parsedEntry, nil
}

func parseStringAnswerLine(entry string) ([]int, error) {
	digits := strings.Split(entry, " ")
	validResult := make([]int, 0, len(digits))
	for _, digit := range digits {
		asInteger, parseErr := strconv.Atoi(digit)
		if parseErr != nil {
			return []int{}, parseErr
		}
		validResult = append(validResult, asInteger)
	}
	return validResult, nil
}

func TestCombinationSum(t *testing.T) {
	stringAnswers := []string{
		"1 1 1 1 1 1 1",
		"1 1 1 1 1 2",
		"1 1 1 1 3",
		"1 1 1 2 2",
		"1 1 2 3",
		"1 2 2 2",
		"1 3 3",
		"2 2 3",
	}

	expectedResult, err := parseStringAnswer(stringAnswers)
	if err != nil {
		t.Errorf("Parsing error! %s\n", err)
	}

	result := combinationSum([]int{1, 2, 3}, 7)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}
}
