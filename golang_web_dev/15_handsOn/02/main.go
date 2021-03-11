package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// $ telnet localhost 8080

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("End of the http request header")
				break
			}
		}

		fmt.Println("code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
