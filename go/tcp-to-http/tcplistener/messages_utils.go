package main

import (
	"io"
	"log"
	"os"
)

func ReadMessages() io.ReadCloser {
	file, err := os.Open("resources/messages.txt")
	if err != nil {
		log.Fatal("error reading messages.txt")
	}
	return file
}
