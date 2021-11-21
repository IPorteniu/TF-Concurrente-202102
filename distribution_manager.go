package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var localhost string

var AZ1 bool = true
var AZ2 bool = true

func sender(ip string, puerto string, data string) {
	ln, err := net.Listen("tcp", ip+":"+puerto)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	con, err := net.Dial("tcp", ip+":"+puerto)
	if err != nil {
		fmt.Println("Error al conectar", err)
	}
	defer con.Close()
	fmt.Fprintln(con, data)
}

func receiver(ip string, puerto string) {
	// receive
	ln, err := net.Listen("tcp", ip+":"+puerto)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		con, err := ln.Accept()
		fmt.Println("Connection accepted", con.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		go senderConnectionHandler(con)
	}

}

func distributionManager(port string, con net.Conn, data string) {
	// Leemos lo que llega de la conexión con los nodos
	// Si la comunicación es por el puerto 9090, entonces se envia a nodo 1 o nodo 2 dependiendo si está ocupado o no
	if port == "9090" {
		// Enviar a nodo disponible
		fmt.Fprintln(con, "prueba completada")
		fmt.Println("Se distribuye")
		//sender(localhost, "9095", data)
		// colocar criterios de distribución
		//sender(localhost, "9096", data)
	} else if port == "9095" || port == "9096" { // Si la comunicación es por los puertos 9095 o 9096, entonces se envia al backend
		// Enviar a backend
		//sender(localhost, "9090", data)
	}
}

func senderConnectionHandler(con net.Conn) {
	defer con.Close()
	// Leemos lo que llega de la conexión con los nodos
	bufferI := bufio.NewReader(con)
	data, _ := bufferI.ReadString('\n')
	// Extraer puerto del local address y distribuir las cargas dependiendo de eso
	_, port, err := net.SplitHostPort(con.LocalAddr().String())
	if err != nil {
		log.Fatal(err)
	}
	distributionManager(port, con, data) // borré la goroutine, sigue funcionando

	fmt.Println(port)
	fmt.Printf(data)
}

func main() {
	//configuracion
	localhost = "192.168.1.90"
	// Escucha en el backend
	go receiver(localhost, "9090")
	// Escucha en nodo 1
	go receiver(localhost, "9095")
	// Escucha en nodo 2
	go receiver(localhost, "9096")

	fmt.Scanf("Enter")

}
