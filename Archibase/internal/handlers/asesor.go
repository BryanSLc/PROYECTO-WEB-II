package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

func (s *Servidor) CrearAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevoAsesor models.Asesor
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevoAsesor)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevoAsesor.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	nuevoAsesor.IDasesor = storage.ConteoAsesores
	storage.ConteoAsesores++
	storage.ListaAsesores = append(storage.ListaAsesores, nuevoAsesor)

	fmt.Println("--> Asesor creado con ID:", nuevoAsesor.IDasesor)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevoAsesor)
}

func (s *Servidor) ObtenerAsesores(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	json.NewEncoder(respuesta).Encode(storage.ListaAsesores)
}

func (s *Servidor) ObtenerAsesorPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for _, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			fmt.Println("--> Asesor encontrado con ID:", id)
			json.NewEncoder(respuesta).Encode(asesor)
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func (s *Servidor) ActualizarAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var asesorActualizado models.Asesor
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&asesorActualizado)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if asesorActualizado.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	asesorActualizado.IDasesor = id
	for i, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			storage.ListaAsesores[i] = asesorActualizado
			fmt.Println("--> Asesor actualizado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje": "Asesor actualizado",
				"asesor":  asesorActualizado,
			})
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}

func (s *Servidor) EliminarAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	for i, asesor := range storage.ListaAsesores {
		if asesor.IDasesor == id {
			storage.ListaAsesores = append(storage.ListaAsesores[:i], storage.ListaAsesores[i+1:]...)
			fmt.Println("--> Asesor eliminado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Asesor eliminado correctamente"})
			return
		}
	}

	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
}
