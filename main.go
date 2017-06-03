package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

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

func main() {
	usuarios = append(usuarios, Usuario{ID: "1", Nombre: "Javier", Apellido: "Coppis", Direccion: &Direccion{Ciudad: "Tandil", Calle: "Tierra del Fuego", NroCalle: 967}})
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", GetUsuariosEndPoint).Methods("GET")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("GET")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("POST")
	router.HandleFunc("/usuarios/{id}", GetUsuariosEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
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
