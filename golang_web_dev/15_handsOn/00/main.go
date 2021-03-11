package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// $ telnet localhost 8080

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
			continue
		}
		fmt.Fprintln(conn, "fprint: hello")
		io.WriteString(conn, "io string: hello")

		conn.Close()
	}
}
