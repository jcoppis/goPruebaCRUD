package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

//Facultad a
type Facultad struct {
	ID       string     `json:"id,omitempty"`
	Nombre   string     `json:"nombre,omitempty"`
	Carreras *[]Carrera `json:"carrera,omitempty"`
}

//Carrera a
type Carrera struct {
	ID       string     `json:"id,omitempty"`
	Nombre   string     `json:"nombre,omitempty"`
	Materias *[]Materia `json:"materia,omitempty"`
}

//Materia a
type Materia struct {
	ID     string `json:"id,omitempty"`
	Nombre string `json:"nombre,omitempty"`
}

type Usuario struct {
	ID        string     `json:"id,omitempty"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Direccion *Direccion `json:"direccion,omitempty"`
}

type Direccion struct {
	Ciudad   string `json:"ciudad,omitempty"`
	Calle    string `json:"calle,omitempty"`
	NroCalle int    `json:"nro_calle,omitempty"`
}

var usuarios []Usuario
var facultades []Facultad
var carreras []Carrera
var materias1, materias2 []Materia

func main() {
	usuarios = append(usuarios, Usuario{ID: "1", Nombre: "Javier", Apellido: "Coppis", Direccion: &Direccion{Ciudad: "Tandil", Calle: "Tierra del Fuego", NroCalle: 967}})

	materias1 = append(materias1, Materia{ID: "1", Nombre: "web 1"})
	materias1 = append(materias1, Materia{ID: "2", Nombre: "web 2"})
	materias2 = append(materias2, Materia{ID: "3", Nombre: "Algoritmos 1"})
	materias2 = append(materias2, Materia{ID: "4", Nombre: "Ingles"})

	carreras = append(carreras, Carrera{ID: "1", Nombre: "tudai", Materias: &materias1})
	carreras = append(carreras, Carrera{ID: "2", Nombre: "ing en sistemas", Materias: &materias2})

	facultades = append(facultades, Facultad{ID: "1", Nombre: "Exactas", Carreras: &carreras})

	router2 := mux.NewRouter()
	router2.HandleFunc("/facultades", GetFacultadesEndPoint).Methods("GET")
	router2.HandleFunc("/facultades/{id}", GetFacultadEndPoint).Methods("GET")
	router2.HandleFunc("/facultades/{id}", CreateFacultadEndPoint).Methods("POST")
	router2.HandleFunc("/facultades/{id}", DeleteFacultadEndPoint).Methods("DELETE")

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", GetUsuariosEndPoint).Methods("GET")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("GET")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("POST")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router2))
}

func GetFacultadesEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(facultades)
}

func GetFacultadEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range facultades {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateFacultadEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var facultad Facultad
	json.NewDecoder(req.Body).Decode(&facultad)
	facultad.ID = params["id"]
	facultades = append(facultades, facultad)
	json.NewEncoder(w).Encode(usuarios)
}

func DeleteFacultadEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range facultades {
		if item.ID == params["id"] {
			facultades = append(facultades[:index], facultades[index+1:]...)
			break
		}
	}
}

func GetUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range usuarios {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func GetUsuariosEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

func CreateUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var usuario Usuario
	json.NewDecoder(req.Body).Decode(&usuario)
	usuario.ID = params["id"]
	usuarios = append(usuarios, usuario)
	json.NewEncoder(w).Encode(usuarios)
}

func DeleteUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range usuarios {
		if item.ID == params["id"] {
			usuarios = append(usuarios[:index], usuarios[index+1:]...)
			break
		}
	}
}
