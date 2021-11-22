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
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	DNI       int     `json:"dni"`
	Edad      float64 `json:"edad"`
	Tipo      float64 `json:"tipo"`
	Actividad float64 `json:"actividad"`
	Insumo    float64 `json:"insumo"`
	Metodo    string  `json:"metodo"`
}

var listaUsuaria []Usuaria

func loadData() {
	listaUsuaria = []Usuaria{}
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
		//fmt.Print("imprimir usuariasJSON")
		json.Unmarshal(cuerpoMsg, &newUsuaria)
		newUsuaria.ID = len(listaUsuaria) + 1
		listaUsuaria = append(listaUsuaria, newUsuaria)
		fmt.Print(listaUsuaria)
		go handle(newUsuaria)
		fmt.Print("salio")
		json.NewEncoder(res).Encode(newUsuaria)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	}
}
func handle(newUsuaria Usuaria) {
	con, _ := net.Dial("tcp", "localhost:9000")
	defer con.Close()
	fmt.Fprintln(con, newUsuaria)
	//r := bufio.NewReader(con)
	//resp, _ := r.ReadString('\n')
	//fmt.Printf("%s", resp)
}

func receiver(ip string, puerto string) {
	// receive
	ln, err := net.Listen("tcp", ip+":"+puerto)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	con, err := ln.Accept()
	fmt.Println("Connection accepted", con.LocalAddr())
	if err != nil {
		log.Fatal(err)
	}
	bufferIn := bufio.NewReader(con)
	mensaje, _ := bufferIn.ReadString('\n')
	fmt.Println(mensaje)
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
	localhost = "localhost"
	remotehost = "localhost"
	go receiver(localhost, "9001")
	loadData()
	Routes()

}
