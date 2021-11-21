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
	con, err := net.Dial("tcp", "192.168.1.90:8000")
	defer con.Close()
	if err != nil {
		fmt.Println("Error al conectar", err)
	}
	fmt.Fprintln(con, listaUsuaria)
}
