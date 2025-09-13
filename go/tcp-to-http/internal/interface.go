package internal

import (
	"bufio"
	"fmt"
	"log"
)

func ServerInterface(buffer *bufio.Reader, callback func([]byte) (int, error)) {
	fmt.Print("> ")
	input, inputErr := buffer.ReadString('\n')

	if inputErr != nil {
		log.Fatalf("input error: %s", inputErr)
	}
	writeSize, writeErr := callback([]byte(input))
	if writeErr != nil {
		log.Fatalf("write error: %s", writeErr)
	} else {
		fmt.Printf("successfully wrote %d bytes to udp\n", writeSize)
	}
}
