package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

func myIp() string {
	ifaces, err := net.Interfaces()
	// Manejador err
	if err != nil {
		log.Print(fmt.Errorf("localAddres: %v \n", err.Error()))
		return "127.0.0.1"
	}

	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "Ethernet") {
			addrs, err := iface.Addrs()
			// Manejador err
			if err != nil {
				log.Print(fmt.Errorf("localAddres: %v \n", err.Error()))
				return "127.0.0.1"
			}

			for _, addr := range addrs {
				switch d := addr.(type) {
				case *net.IPNet:
					if strings.HasPrefix(d.IP.String(), "192") {
						return d.IP.String()
					}
				}
			}
		}
	}
	return "127.0.0.1"
}

func main() {
	cargardata()
	con, err := net.Dial("tcp", myIp()+":9090")
	defer con.Close()
	if err != nil {
		fmt.Println("Error al conectar", err)
	}
	fmt.Fprintln(con, listaUsuaria)
	r := bufio.NewReader(con)
	resp, _ := r.ReadString('\n')
	fmt.Printf("Respuesta: %s", resp)
}
