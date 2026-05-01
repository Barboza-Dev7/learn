package main

import (
	"log"
	"net"
	"encoding/hex"
	"io"
)

func hanldeConnection(conn net.Conn){
	defer conn.Close()
	log.Println("Conectada: ", conn.RemoteAddr())

	buf := make([]byte, 128)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			log.Println(err)
			return
		}
		if io.EOF == nil {
			return
		}
		if n < 6 {
			continue
		}
		raw := buf[:n]
		log.Println(hex.EncodeToString(raw))
	}
}

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Escuchando ...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go hanldeConnection(conn)
	}
}