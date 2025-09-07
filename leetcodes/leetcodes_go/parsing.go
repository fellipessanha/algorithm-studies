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

// func sayType[T any](entry T, outputType Type) {
// 	typeKind := reflect.TypeOf(entry).Kind()
// 	fmt.Println(typeKind)
// 	fmt.Println(outputType)

// 	switch typeKind {
// 	case reflect.Int:
// 		fmt.Println("oi")
// 	case reflect.String:
// 		fmt.Println("eai")
// 	case reflect.Slice:
// 		fmt.Println("tchaus")
// 	default:
// 		fmt.Println("nao foi")
// 	}
// }

// func main() {
// 	sayType("oi", int)
// 	sayType(123, string)
// 	sayType([]int{3, 4}, map[int]int)
// }
