package main

import (
	"net"
	"bytes"
	"io"
	"fmt"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil{
		panic(err)
	}
	conn, err := listener.Accept()
	conn, err := net.Dial("tcp", "127.0.0.1:8085")
	var dataBuffer bytes.Buffer
	b:=make([]byte, 10)
	for {
		n, err := conn.Read(b)
		if err!=nil {
			if err==io.EOF{
				fmt.Println("The connection is closed.")
				conn.Close()
			} else {
				fmt.Printf("Read Error: %s\n", err)
			}
			break
		}
		dataBuffer.Write(b[:n])
	}
	reader := bufio.NewReader(conn)
	line, err:=reader.ReadBytes('\n')
}
