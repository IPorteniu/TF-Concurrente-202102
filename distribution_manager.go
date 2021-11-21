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
		fmt.Println("Connection accepted", con.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		go connectionHandler(con)
	}

}

func connectionHandler(con net.Conn) {
	defer con.Close()
	// Leemos lo que llega de la conexión con los nodos
	bufferI := bufio.NewReader(con)
	data, _ := bufferI.ReadString('\n')
	fmt.Println(con.LocalAddr().Network())

	fmt.Printf(data)
}

func main() {
	//configuracion

	// Escucha en el backend
	receiver(localhost, "9090")
	// Escucha en nodo 1
	go receiver(localhost, "9095")
	// Escucha en nodo 2
	go receiver(localhost, "9096")

}
