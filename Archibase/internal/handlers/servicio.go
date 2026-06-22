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

type Servidor struct {
	Almacen *storage.SQLiteStorage
}

func NuevoServidor(a *storage.SQLiteStorage) *Servidor {
	return &Servidor{Almacen: a}
}

func (s *Servidor) CrearServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevoServicio models.Servicio
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevoServicio)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevoServicio.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	servicioCreado := s.Almacen.CrearServicio(nuevoServicio)
	fmt.Println("--> Servicio creado con ID:", servicioCreado.ID)

	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(servicioCreado)
}

func (s *Servidor) ObtenerServicios(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los servicios")
	json.NewEncoder(respuesta).Encode(s.Almacen.ListarServicios())
}

func (s *Servidor) ObtenerServicioPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	servicio, encontrada := s.Almacen.BuscarServicioPorID(id)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
		return
	}

	fmt.Println("--> Servicio encontrado con ID:", id)
	json.NewEncoder(respuesta).Encode(servicio)
}

func (s *Servidor) ActualizarServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var servicioActualizado models.Servicio
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&servicioActualizado)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if servicioActualizado.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	servicio, encontrada := s.Almacen.ActualizarServicio(id, servicioActualizado)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
		return
	}

	fmt.Println("--> Servicio actualizado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje":  "Servicio actualizado",
		"servicio": servicio,
	})
}

func (s *Servidor) EliminarServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if !s.Almacen.EliminarServicio(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
		return
	}

	fmt.Println("--> Servicio eliminado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Servicio eliminado correctamente"})
}
