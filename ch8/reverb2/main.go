package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second) // read multiple inputs concurrently
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\t%s\n\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n\t", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n", strings.ToLower(shout))
}
