package leetcodesgo

import (
	"strconv"
	"strings"
)

// Element is an interface that can represent either a string or an integer
type Element interface{ string | int}

func parseGenericArray[T any](
	entry string,
	splittings string,
	trimmings string,
	unitary_parse_callback func(string) (T, error),
) ([]T, error) {
	unitary_elements := strings.Split(strings.Trim(entry, trimmings), splittings)
	parsed_entries := make([]T, 0, len(unitary_elements))

	for _, unitary_element := range unitary_elements {
		parsed_entry, parseErr := unitary_parse_callback(unitary_element)
		if parseErr != nil {
			return []T{}, parseErr
		}
		parsed_entries = append(parsed_entries, parsed_entry)
	}
	return parsed_entries, nil
}

func parseStringArrayArray(entry string) ([][]string, error) {
	return parseGenericArray(entry, "],[", "[]", parseStringArray)
}

func parseStringArrayArrayArray(entry string) ([][][]string, error) {
	return parseGenericArray(entry, "]],[[", "[]", parseStringArrayArray)
}

func parseStringArray(entry string) ([]string, error) {
	return parseGenericArray(entry, ",", "\"[]", func(entry string) (string, error) { return entry, nil })
}

func parseIntegerArrayArray(entry string) ([][]int, error) {
	return parseGenericArray(entry, "],[", "[]", parseIntegerArray)
}

func parseIntegerArray(entry string) ([]int, error) {
	return parseGenericArray(entry, ",", "[]", strconv.Atoi)
}

func parseBinaryNodeList(entry string) (*TreeNode, error) {
	values, _ := parseStringArray(entry)
	return buildBinaryTreeFromStringArray(values, 0)
}

func parseListNodeList(entry string) (*ListNode, error) {
	array, err := parseIntegerArray(entry)
	if err != nil {
		return nil, err
	}

	return buildListNode(array), nil
}