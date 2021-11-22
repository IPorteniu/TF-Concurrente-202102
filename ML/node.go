package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

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

type DataSet struct {
	Data   [][]interface{}
	Labels []string
}

func readDataSet() [][]string {
	// Obtener el dataset desde github
	metodoMatrix := [][]string{}
	url := "https://github.com/IPorteniu/TF-Concurrente-202102/raw/main/Data/DAT%20PlaniFamiliar_01_Metodo.csv"
	dataset, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer dataset.Body.Close()

	// Maneja la codificación del archivo si es que hubiera
	br := bufio.NewReader(dataset.Body)
	r, _, err := br.ReadRune()
	if err != nil {
		panic(err)
	}
	if r != '\uFEFF' {
		br.UnreadRune()
	}

	// Leer el dataset
	reader := csv.NewReader(br)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		metodoMatrix = append(metodoMatrix, record)
	}

	return metodoMatrix
}

func loadData() DataSet {

	// Cargar el DataSet desde su CSV
	data := readDataSet()
	ds := DataSet{}

	// Inicializar la usuaria Struct para llenarlo con datos
	usuaria := Usuaria{}

	// Almacenar los datos en las estructuras
	for i, metodos := range data {
		// Drop de la primera fila (titles)
		if i == 0 {
			continue
		}

		temp := make([]interface{}, 0)
		// Convertimos los datos necesarios a floats para poder añadirlos
		for j, value := range metodos[:] {

			if j == 6 {
				switch value {
				case "12 a - 17 a":
					usuaria.Edad = 14.5
				case "18 a - 29 a":
					usuaria.Edad = 23.5
				case "30 a - 59 a":
					usuaria.Edad = 44.5
				case "> 60 a":
					usuaria.Edad = 65.0
				}
				temp = append(temp, usuaria.Edad)
			} else if j == 8 {
				// METODO
				usuaria.Metodo = value
			} else if j == 9 {
				// Si son Nuevas = 0 y si son Continuadoras = 1
				switch value {
				case "NUEVAS":
					usuaria.Tipo = 0.0
				case "CONTINUADORAS":
					usuaria.Tipo = 1.0
				}
				// TIPO DE USUARIA
				temp = append(temp, usuaria.Tipo)
			} else if j == 10 {
				parsedValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(err)
				}
				// ACTIVIDAD
				usuaria.Actividad = parsedValue
				temp = append(temp, usuaria.Actividad)
			} else if j == 11 {
				parsedValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(err)
				}
				// INSUMO
				usuaria.Insumo = parsedValue
				temp = append(temp, usuaria.Insumo)
			}

		}
		// Filtramos todas las filas que contengan MELA ya que no es un Metodo anticonceptivo que se pueda recomendar normalmente
		if metodos[7] != "MELA" {

			// Añadir los datos al DataSet struct ahora convertidos
			ds.Data = append(ds.Data, temp)
			ds.Labels = append(ds.Labels, metodos[8])
		}
	}
	return ds
}

func extractFeatures(u []Usuaria) [][]interface{} {

	features := [][]interface{}{}
	for i := range u {

		featureData := []interface{}{u[i].Edad, u[i].Tipo, u[i].Actividad, u[i].Insumo}
		features = append(features, featureData)
	}

	return features
}

func main() {
	// ENTRENAMIENTO DATASET
	ds := loadData()
	fmt.Println(len(ds.Data))

	forest := TrainForest(ds.Data, ds.Labels, len(ds.Data)/10, len(ds.Data[0]), 5)

	// DATOS DE USUARIA
	usuaria1 := Usuaria{ID: 20, Edad: 5., Tipo: 3.5, Actividad: 1.4, Insumo: 0.2} //Setosa
	usuaria2 := Usuaria{ID: 21, Edad: 7, Tipo: 3.2, Actividad: 4.7, Insumo: 1.4}  //Versicolor
	usuaria3 := Usuaria{ID: 22, Edad: 6.3, Tipo: 3.3, Actividad: 6, Insumo: 2.5}  // Virginica
	usuariasJSON := []Usuaria{usuaria1, usuaria2, usuaria3}
	usuariaX := [][]interface{}{}
	for i, _ := range usuariasJSON {
		usuariaI := []interface{}{usuariasJSON[i].Edad, usuariasJSON[i].Tipo, usuariasJSON[i].Actividad, usuariasJSON[i].Insumo}
		usuariaX = append(usuariaX, usuariaI)
	}

	fmt.Println("USUARIAS")
	fmt.Println(usuariasJSON)
	fmt.Println("EXTRACTING FEATS")
	features := extractFeatures(usuariasJSON)
	fmt.Println(features)

	// OUTPUT

	fmt.Println("USUARIAS CON PREDICCION")
	var output string
	for i := 0; i < len(features); i++ {
		output = forest.Predicate(features[i])
		usuariasJSON[i].Metodo = output
	}
	fmt.Println(usuariasJSON)
}
