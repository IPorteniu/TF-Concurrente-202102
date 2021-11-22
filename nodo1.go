package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "192.168.1.90:9095")
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handle(con) // podemos atender miles de clienes concurrentemente!
	}
}

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)

	msg, _ := r.ReadString('\n')
	fmt.Printf("Recibido: %s", msg)
	fmt.Fprintln(con, "Se recibió de nodo 1") // acá va la respuesta de fabrizio

}
