package main

import (
	"fmt"
	"net"
)

type Usuaria struct {
	Nombre string
	Edad   int
	Dni    int
}

var listaUsuaria []Usuaria

func cargardata() {
	listaUsuaria = []Usuaria{
		{"Ivana", 23, 72837245},
		{"Yvana", 19, 87236732},
		{"Sebastiana", 20, 76546378}}
}

func main() {
	cargardata()
	con, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error al conectar", err)
	}
	defer con.Close()
	fmt.Fprintln(con, listaUsuaria)
}
