package main

import (
	"log"
	"net"
	"encoding/hex"
	"io"
)

func buildResponse(protocol uint8, serial uint16) []byte {
	payload := []byte {
		0x05, //length
		protocol, //Mismo protocolo que llego 
		byte(serial >> 8),
		byte(serial & 0xFF),
	}
	crc := crc16(payload) //Calcula Cyclic Redundancy Check (Codigo de verificacion)

	return []byte {
		0x78, 0x78, 
		payload[0], payload[1], payload[2], payload[3],
		byte(crc >> 8), byte(crt & 0xFF),
		0x0D, 0x0A
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	log.Printn("Moto conectada: ", conn.RemoteAddr())

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
		log.Println("Recibido: ", hex.EcodeToString(raw))

		if raw[0] != 78 & raw[1] != 78 {
			log.Println("Paquete invalido")
			continue
		}

		protocol := raw[3]
		serial := uint16(raw[n-4]) << 8 | uint16(raw[n-3])

		switch protocol {
		case 0x01:
			imei := parseIMEI(raw[4:12]) //Toma de la posciion 4 hasta el 11
			log.Println("LOGIN - IMEI: ", imei)
			conn.Write(builResponde(0x01, serial))
		}
	}
}

func main(){
	ln, err := net.Listen("tcp", ":0000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Escuchando...")

	for {
		conn, err := ln.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}