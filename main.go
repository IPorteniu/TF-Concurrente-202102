package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var localhost string
var remotehost string

type Usuaria struct {
	ID     int
	Nombre string
	Años   int
	Dni    int
}

var listaUsuaria []Usuaria

func loadData() {
	listaUsuaria = []Usuaria{
		{1, "Ivana", 23, 72837245},
		{2, "Yvana", 19, 87236732},
		{3, "Sebastiana", 20, 76546378}}
}

func Routes() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/dataset", MuestraDataSet)
	mux.HandleFunc("/api/agregar", agregarUsuaria)
	log.Fatal(http.ListenAndServe(":9080", mux))
}

func MuestraDataSet(res http.ResponseWriter, req *http.Request) {
	log.Println("llamada al endpoint /dataset")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(listaUsuaria)
}
func agregarUsuaria(res http.ResponseWriter, req *http.Request) {
	var newUsuaria Usuaria
	if req.Method == "POST" {
		log.Println("Ingreso al metodo agregar")
		cuerpoMsg, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Error interno al leer el body", http.StatusInternalServerError)
		}
		fmt.Print("imprimir usuariasJSON")
		json.Unmarshal(cuerpoMsg, &newUsuaria)
		newUsuaria.ID = len(listaUsuaria) + 1
		listaUsuaria = append(listaUsuaria, newUsuaria)
		fmt.Print(listaUsuaria)
		handle(newUsuaria)
		fmt.Print("salio")
		json.NewEncoder(res).Encode(newUsuaria)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
func handle(newUsuaria Usuaria) {
	con, _ := net.Dial("tcp", remotehost+":9090")
	defer con.Close()
	r := bufio.NewReader(con)
	fmt.Fprintln(con, newUsuaria)
	resp, _ := r.ReadString('\n')
	fmt.Printf("%s", resp)
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
	fmt.Println(port)
	fmt.Printf(data)
}
func main() {
	localhost = "192.168.0.90"
	remotehost = "201.230.178.131"
	go receiver(localhost, "9090")
	loadData()
	Routes()

}
