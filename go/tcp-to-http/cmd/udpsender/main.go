package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	address, addressErr := net.ResolveUDPAddr("udp", "localhost:42069")
	if addressErr != nil {
		log.Fatalf("erro resolvendo endereço: %s", addressErr)
	} else {
		fmt.Printf("endereço resolvido: %s\n", address)
	}
	connection, connectionErr := net.DialUDP("udp", nil, address)
	if connectionErr != nil {
		log.Fatalf("erro preparando conexão %s", connectionErr)
	} else {
		log.Printf("conexão pronta: %v\n", connection)
		defer connection.Close()
	}

	buffer := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, inputErr := buffer.ReadString('\n')

		if inputErr != nil {
			log.Fatalf("input error: %s", inputErr)
		}
		writeSize, writeErr := connection.Write([]byte(input))
		if writeErr != nil {
			log.Fatalf("write error: %s", writeErr)
		} else {
			fmt.Printf("successfully wrote %d bytes to udp\n", writeSize)
		}

	}
}
