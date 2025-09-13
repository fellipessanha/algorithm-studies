package main

import (
	"fmt"
	"log"
	"net"
)

func Tcp() {
	tcp, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("conexão não rolou")
	}

	connectionIO, tcpErr := tcp.Accept()
	if tcpErr != nil {
		log.Fatal("conexão não aceitou")
	}

	lines := getLinesChannel(connectionIO)
	for line := range lines {
		fmt.Println(line)
	}
}
