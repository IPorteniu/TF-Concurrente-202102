package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var remotehost string

type Usuaria struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	DNI       int     `json:"dni"`
	Edad      float64 `json:"edad"`
	Tipo      float64 `json:"tipo"`
	Actividad float64 `json:"actividad"`
	Insumo    float64 `json:"insumo"`
	Metodo    string  `json:"metodo"`
}

func main() {

	usuaria1 := Usuaria{ID: 20, Edad: 5, Tipo: 3.5, Actividad: 1.4, Insumo: 0.2}
	for {
		time.Sleep(2 * time.Second)
		send(usuaria1)
	}
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
func send(user Usuaria) {
	//Nodo a cual mandar la usuaria
	conn, _ := net.Dial("tcp", myIp()+":9090")
	defer conn.Close()
	// Codificar JSON
	bytesMsg, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	// Enviar mensaje serializado en string
	fmt.Fprintln(conn, string(bytesMsg))
	r := bufio.NewReader(conn)
	resp, _ := r.ReadString('\n')
	fmt.Println(resp)
}
