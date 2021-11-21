package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:9080")

	if err != nil {
		fmt.Println("Error de conexión - Felipe")
		os.Exit(1)
	}
	defer ln.Close()

	con, _ := ln.Accept()

	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	fmt.Println(msg)
}
