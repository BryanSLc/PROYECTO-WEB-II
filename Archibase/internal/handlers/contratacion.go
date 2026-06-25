package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearContratacion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevaContratacion models.Contratacion
	if err := json.NewDecoder(peticion.Body).Decode(&nuevaContratacion); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	creada, err := s.ContratacionSvc.Crear(nuevaContratacion)
	if err != nil {
		responderError(respuesta, err)
		return
	}

	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(creada)
}

func (s *Servidor) ObtenerContrataciones(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respuesta).Encode(s.ContratacionSvc.Listar())
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

	contratacion, err := s.ContratacionSvc.BuscarPorID(id)
	if err != nil {
		responderError(respuesta, err)
		return
	}

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

	var datos models.Contratacion
	if err := json.NewDecoder(peticion.Body).Decode(&datos); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	actualizada, err := s.ContratacionSvc.Actualizar(id, datos)
	if err != nil {
		responderError(respuesta, err)
		return
	}

	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje":      "Contratacion actualizada",
		"contratacion": actualizada,
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

	if err := s.ContratacionSvc.Eliminar(id); err != nil {
		responderError(respuesta, err)
		return
	}

	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Contratacion eliminada correctamente"})
}
