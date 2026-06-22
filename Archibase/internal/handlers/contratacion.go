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

func (s *Servidor) CrearContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevaContratacion models.Contratacion
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaContratacion)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevaContratacion.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	contratacionCreada := s.Almacen.CrearContratacion(nuevaContratacion)
	fmt.Println("--> Contratacion creada con ID:", contratacionCreada.ID)

	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(contratacionCreada)
}

func (s *Servidor) ObtenerContrataciones(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las contrataciones")
	json.NewEncoder(respuesta).Encode(s.Almacen.ListarContrataciones())
}

func (s *Servidor) ObtenerContratacionPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	contratacion, encontrada := s.Almacen.BuscarContratacionPorID(id)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
		return
	}

	fmt.Println("--> Contratacion encontrada con ID:", id)
	json.NewEncoder(respuesta).Encode(contratacion)
}

func (s *Servidor) ActualizarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var contratacionActualizada models.Contratacion
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&contratacionActualizada)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if contratacionActualizada.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	contratacion, encontrada := s.Almacen.ActualizarContratacion(id, contratacionActualizada)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
		return
	}

	fmt.Println("--> Contratacion actualizada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje":      "Contratacion actualizada",
		"contratacion": contratacion,
	})
}

func (s *Servidor) EliminarContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if !s.Almacen.EliminarContratacion(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Contratacion no encontrada"})
		return
	}

	fmt.Println("--> Contratacion eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Contratacion eliminada correctamente"})
}
