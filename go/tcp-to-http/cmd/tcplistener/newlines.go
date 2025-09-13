package main

import (
	"fmt"
	"strings"
)

func ParseNewLines() []string {
	var response []string
	file := ReadMessages()
	text := ""
	for {
		bytes := make([]byte, 8)
		_, err := file.Read(bytes)
		if err != nil {
			break
		}
		text += string(bytes)

		if pos := strings.IndexByte(text, '\n'); pos != -1 {
			response = append(response, text[:pos])
			text = text[pos+1:]
		}
	}
	return append(response, text)
}

func Newlines() {
	for _, line := range ParseNewLines() {
		fmt.Printf("read: %s\n", line)
	}

}
