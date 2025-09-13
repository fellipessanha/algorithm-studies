package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"tcp-to-http/internal"
)

func main() {
	connection := getUDPConnection()

	buffer := bufio.NewReader(os.Stdin)
	for {
		internal.ServerInterface(buffer, connection.Write)
	}
}

func getUDPConnection() *net.UDPConn {
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
	return connection
}
