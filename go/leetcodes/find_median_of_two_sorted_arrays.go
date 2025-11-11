package leetcodesgo

import (
	"sort"
)

func findMedianSortedArrays(l1 []int, l2 []int) float64 {
	for _, el := range l2 {
		l1 = append(l1, el)
	}
	sort.Ints(l1)

	halfway := len(l1) / 2
	if len(l1)%2 == 0 {
		return float64(l1[halfway]+l1[halfway-1]) / 2.0
	}

	return float64(l1[halfway])
}
