package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func printReader(r io.ReadCloser) {
	defer r.Close()
	buf := bytes.NewBuffer(make([]byte, 1024))
	l, err := buf.ReadFrom(r)
	fmt.Print(buf)
	for err != nil && l > 0 {
		l, err = buf.ReadFrom(r)
		fmt.Print(buf)
	}
}

func main() {
	ss, err := net.Listen("tcp", ":6288")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ss.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go printReader(conn)
	}
}
