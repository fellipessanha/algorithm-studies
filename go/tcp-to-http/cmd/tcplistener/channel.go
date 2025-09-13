package main

import (
	"fmt"
	"io"
	"strings"
)

func GetLinesFromChannel() {
	lines := getLinesChannel(ReadMessages())
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string, 1)
	go func(file io.ReadCloser) {
		defer file.Close()
		defer close(channel)
		text := ""
		for {
			bytes := make([]byte, 8)
			_, err := file.Read(bytes)
			if err != nil {
				break
			}
			text += string(bytes)

			if pos := strings.IndexByte(text, '\n'); pos != -1 {
				channel <- text[:pos]
				text = text[pos+1:]
			}
		}
		channel <- text
	}(f)

	return channel
}
