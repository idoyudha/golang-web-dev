package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	io.WriteString(c, "I see you connected")
}
