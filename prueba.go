package main

import (
	"log"
	"net"
	"encoding/hex"
	"encoding/binary"
	"io"
	"math"
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
		if err == io.EOF {
			log.Println("Desconectado")
			return
		}
		if n < 6 {
			continue
		}

		raw := buf[:n]
		protocol := raw[3]
		//log.Println(hex.EncodeToString(raw))

		switch protocol {
		case 0x01: 
			conn.Write(raw)
		case 0x12:
			latRaw := binary.BigEndian.Uint32(raw[11:15])
			lonRaw := binary.BigEndian.Uint32(raw[15:19])

			lat := float64(latRaw) / 1800000.0
			lon := float64(lonRaw) / 1800000.0

			latClean := math.Round(lat*1e6) / 1e6
			lonClean := math.Round(lon*1e6) / 1e6

			log.Println("lat: ", latClean, " - ", "lon: ", lonClean)
		}
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