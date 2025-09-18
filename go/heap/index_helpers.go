package heap

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftIndex(i int) int {
	return (i * 2) + 1
}

func rightIndex(i int) int {
	return (i + 1) * 2
}
