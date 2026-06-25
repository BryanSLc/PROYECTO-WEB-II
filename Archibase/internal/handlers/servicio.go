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

func (s *Servidor) CrearServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevoServicio models.Servicio
	if err := json.NewDecoder(r.Body).Decode(&nuevoServicio); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nuevoServicio.IDservicio = storage.ConteoServicios
	storage.ConteoServicios++
	storage.ListaServicios = append(storage.ListaServicios, nuevoServicio)

	fmt.Println("--> Servicio creado con ID:", nuevoServicio.IDservicio)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoServicio)
}

func (s *Servidor) ObtenerServicios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los servicios")
	json.NewEncoder(w).Encode(storage.ListaServicios)
}

func (s *Servidor) ObtenerServicioPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, servicio := range storage.ListaServicios {
		if servicio.IDservicio == id {
			fmt.Println("--> Servicio encontrado con ID:", id)
			json.NewEncoder(w).Encode(servicio)
			return
		}
	}

	fmt.Println("--> Servicio no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}

func (s *Servidor) ActualizarServicio(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var servicioActualizado models.Servicio
	if err := json.NewDecoder(r.Body).Decode(&servicioActualizado); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	servicioActualizado.IDservicio = id

	for i, servicio := range storage.ListaServicios {
		if servicio.IDservicio == id {
			storage.ListaServicios[i] = servicioActualizado
			fmt.Println("--> Servicio actualizado con ID:", id)
			json.NewEncoder(w).Encode(servicioActualizado)
			return
		}
	}

	fmt.Println("--> Servicio no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}

func (s *Servidor) EliminarServicio(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, servicio := range storage.ListaServicios {
		if servicio.IDservicio == id {
			storage.ListaServicios = append(storage.ListaServicios[:i], storage.ListaServicios[i+1:]...)
			fmt.Println("--> Servicio eliminado con ID:", id)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	fmt.Println("--> Servicio no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}
