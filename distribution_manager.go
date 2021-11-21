package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var localhost string

func sender(ip string, puerto string) {
	for {
		// send
	}

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
		fmt.Println("Connection accepted", con.LocalAddr())
		if err != nil {
			log.Fatal(err)
		}
		go connectionHandler(con)
	}

}

func distributionManager(port string, con net.Conn) {
	// Leemos lo que llega de la conexión con los nodos
	// Si la comunicación es por el puerto 9090, entonces se envia a nodo 1 o nodo 2 dependiendo si está ocupado o no
	// Si la comunicación es por los puertos 9095 o 9096, entonces se envia al backend
}

func connectionHandler(con net.Conn) {
	defer con.Close()
	// Leemos lo que llega de la conexión con los nodos
	bufferI := bufio.NewReader(con)
	data, _ := bufferI.ReadString('\n')
	// Extraer puerto del local address y distribuir las cargas dependiendo de eso
	_, port, err := net.SplitHostPort(con.LocalAddr().String())
	if err != nil {
		log.Fatal(err)
	}
	go distributionManager(port, con)
	fmt.Println(port)
	fmt.Printf(data)
}

func main() {
	//configuracion

	// Escucha en el backend
	go receiver(localhost, "9090")
	// Escucha en nodo 1
	go receiver(localhost, "9095")
	// Escucha en nodo 2
	go receiver(localhost, "9096")

	fmt.Scanf("Enter")

}
