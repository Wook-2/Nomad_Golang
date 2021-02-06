package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How are you?")
		fmt.Fprintf(conn, "%v", "Good!")

		conn.Close()
	}
}
