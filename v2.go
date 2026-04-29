//Manejarla he Imprimir datos crudos de una conexion tcp
//Crear una conexion tcp
package main

import(
	"encoding/hex" //Convertir datos binarios a hexadecimal
	"net"
	"log"
	"io" //Leer/escribir datos (streams)
)
func hanldeConnection(conn net.Conn){
	defer conn.Close()
	log.Println("Moto conectada: ", conn.RemoteAddr())

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF { //“Fin del flujo de datos” (End Of File)
			log.Println("Moto desconectada")
			return
		}
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		log.Println("Recibido: ", hex.EncodeToString(buf[:n]))
		//“Toma desde el inicio del slice hasta la posición n (sin incluir n)”
	}
}

func main(){
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Escuchando...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
} 