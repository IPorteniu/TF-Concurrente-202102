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
	con, _ := net.Dial("tcp", "localhost:9090")
	defer con.Close()
	fmt.Fprintln(con, listaUsuaria)
}
