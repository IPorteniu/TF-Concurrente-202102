package main

import (
	"bufio"
	"fmt"
	"net"
)

var temp = 0

func main() {
	escuchar()

}
func escuchar() {

	ln, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		fmt.Println("Error de conexión - Felipe")
	}

	defer ln.Close()

	for {
		con, _ := ln.Accept()

		defer con.Close()

		bufferIn := bufio.NewReader(con)
		msg, _ := bufferIn.ReadString('\n')
		fmt.Println(msg)
		sender(msg)

	}

}
func sender(msg string) {
	temp = temp + 1
	msg = "mensaje recibido"
	con1, _ := net.Dial("tcp", "localhost:9001")
	defer con1.Close()
	fmt.Fprintln(con1, msg, temp)
}
