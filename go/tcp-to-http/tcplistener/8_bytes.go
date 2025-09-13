package main

import (
	"fmt"
)

func EightBytes() {
	file := ReadMessages()
	for {
		bytes := make([]byte, 8)
		size, err := file.Read(bytes)
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(bytes[:size]))
	}
}
