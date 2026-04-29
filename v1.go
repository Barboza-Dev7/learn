//Crear una conexion tcp
package main

import(
	"fmt"
	"net"
	"log"
)

func main(){
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Escuchando...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Moto conectada", conn.RemoteAddr())
	}
}