package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearAsesor(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevoAsesor models.Asesor
	lector := json.NewDecoder(peticion.Body)
	if err := lector.Decode(&nuevoAsesor); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	asesorCreado, svcErr := s.AsesorSvc.Crear(nuevoAsesor)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Asesor creado con ID:", asesorCreado.IDasesor)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(asesorCreado)
}

func (s *Servidor) ObtenerAsesores(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los asesores")
	asesores, svcErr := s.AsesorSvc.Listar(), error(nil)
	_ = svcErr
	json.NewEncoder(respuesta).Encode(asesores)
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

	asesor, svcErr := s.AsesorSvc.BuscarPorID(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
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

	asesorActualizado.IDasesor = id
	asesor, svcErr := s.AsesorSvc.Actualizar(id, asesorActualizado)
	if svcErr != nil {
		responderError(respuesta, svcErr)
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

	if svcErr := s.AsesorSvc.Eliminar(id); svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Asesor eliminado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Asesor eliminado correctamente"})
}
