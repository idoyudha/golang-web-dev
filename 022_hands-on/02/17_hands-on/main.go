package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
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
	var i int
	var method, url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// Package strings implements simple functions to manipulate UTF-8 encoded strings.
		if i == 0 {
			line := strings.Fields(ln)
			method = line[0]
			url = line[1]
			fmt.Println("METHOD:", method)
			fmt.Println("URI:", url)
			fmt.Printf("i = %v line -> %v", i, line)
		}
		if ln == "" {
			break
		}
		i++
	}
	switch {
	case method == "GET" && url == "/":
		handleHomepage(c)
	case method == "GET" && url == "/apply":
		handleApplyGet(c)
	case method == "POST" && url == "/apply":
		handleApplyPost(c)
	default:
		handleDefault(c)
	}
}

func write(c net.Conn, b string) {
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(b))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, b)
}

func handleHomepage(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>Homepage</h1>
			<a href="/">homepage</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	write(c, body)
}

func handleDefault(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>default</h1>
		</body>
		</html>
	`
	write(c, body)
}

func handleApplyGet(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>Apply GET</h1>
			<a href="/">homepage</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="text" name="apply">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	write(c, body)
}

func handleApplyPost(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>Apply POST</h1>
			<a href="/">homepage</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	write(c, body)
}
