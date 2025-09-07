package leetcodesgo

import (
	"strconv"
	"strings"
)

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

func parseIntegerArrayArray(entry string) ([][]int, error) {
	return parseGenericArray(entry, "],[", "[]", parseIntegerArray)
}

func parseIntegerArray(entry string) ([]int, error) {
	return parseGenericArray(entry, ",", "[]", strconv.Atoi)
}

func parseListNodeList(entry string) (*ListNode, error) {
	array, err := parseIntegerArray(entry)
	if err != nil {
		return nil, err
	}

	return buildListNode(array), nil
}
