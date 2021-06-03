package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	// Init server on port 8000
	ln, _ := net.Listen("tcp", ":8000")

	// accept a connection
	conn, _ := ln.Accept()

	// read bytes from conn accepted
	bytes, _ := bufio.NewReader(conn).ReadBytes('\n')

	// print bytes data
	fmt.Println(bytes) // [104 101 108 108 111 10]

	// print string data
	fmt.Println(string(bytes)) // hello
}
