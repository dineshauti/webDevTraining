package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

func main() {

	li, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")

		conn.Close()
	}
}
