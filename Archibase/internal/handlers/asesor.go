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

	asesorCreado := s.Almacen.CrearAsesor(nuevoAsesor)
	fmt.Println("--> Asesor creado con ID:", asesorCreado.IDasesor)

	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(asesorCreado)
}

func (s *Servidor) ObtenerAsesores(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	json.NewEncoder(respuesta).Encode(s.Almacen.ListarAsesores())
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

	asesor, encontrada := s.Almacen.BuscarAsesorPorID(id)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
		return
	}

	fmt.Println("--> Asesor encontrado con ID:", id)
	json.NewEncoder(respuesta).Encode(asesor)
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

	asesor, encontrada := s.Almacen.ActualizarAsesor(id, asesorActualizado)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
		return
	}

	fmt.Println("--> Asesor actualizado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje": "Asesor actualizado",
		"asesor":  asesor,
	})
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

	if !s.Almacen.EliminarAsesor(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Asesor no encontrado"})
		return
	}

	fmt.Println("--> Asesor eliminado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Asesor eliminado correctamente"})
}
