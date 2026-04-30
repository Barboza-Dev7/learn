package main

import (
	"log"
	"net"
	"io"
	"encoding/hex"
)

func handleConnection(conn net.Conn){
	defer conn.Close()
	log.Println("Conectado: ", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
	num, err := conn.Read(buf)

	if err == io.EOF {
		log.Println("Desconectado: ", conn.RemoteAddr())
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	raw := buf[:num]
	protocol := raw[3]

	switch protocol {
	case 0x01:
		log.Println(hex.EncodeToString(raw))
		conn.Write(raw)	
	case 0x12:
		data := hex.EncodeToString(raw)
		log.Println(data)
	}
	}
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