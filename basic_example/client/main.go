package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	fmt.Fprintf(conn, "hello\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)
}
