package main

import (
	"log"
	"net"
	"io"
	"encoding/hex"
)

func handleConnection(conn net.Conn){
	log.Println("Conectado: ", conn.RemoteAddr())

	buf := make([]byte, 1024)
	data, err := conn.Read(buf)
	raw := buf[:data]

	if err == io.EOF {
		log.Println("Desconectado: ", conn.RemoteAddr())
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(hex.EncodeToString(raw))
}

func main(){
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Escuchando...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}