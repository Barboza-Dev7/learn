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
		n, err := conn.Read(buf)
		if err == io.EOF {
			log.Println("Moto desconectada")
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
		raw := buf[:n]
		log.Println("Recibido: ", hex.EncodeToString(raw))
	}
}

func main(){
	listener, err := net.Listen()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Escuchando...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}