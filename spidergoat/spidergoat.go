package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: spidergoat <port>")
		return
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("usage: spidergoat <port>")
	}
	server(port)
}

func server(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	accept, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}

	bytes := make([]byte, 1024)
	_, err = accept.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
