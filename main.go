package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("can't start the server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("started server on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can't accept this connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
