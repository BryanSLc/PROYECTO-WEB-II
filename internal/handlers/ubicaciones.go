package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/service"
)

func (s *Servidor) CrearUbicacion(w http.ResponseWriter, r *http.Request) {
	var nuevaUbicacion models.Ubicacion
	err := json.NewDecoder(r.Body).Decode(&nuevaUbicacion)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	ubicacionCreada, err := s.UbicacionService.Crear(nuevaUbicacion)
	if err != nil {
		if errors.Is(err, service.ErrProvinciaUbicacionObligatoria) || errors.Is(err, service.ErrCiudadUbicacionObligatoria) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Ubicacion creada con ID:", ubicacionCreada.ID)
	RespondJSON(w, http.StatusCreated, ubicacionCreada)
}

func (s *Servidor) ObtenerUbicaciones(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-Obteniendo todas las ubicaciones")
	RespondJSON(w, http.StatusOK, s.UbicacionService.Listar())
}

func (s *Servidor) ObtenerUbicacionPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	ubicacion, err := s.UbicacionService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrUbicacionNoEncontrada) {
			fmt.Println("Ubicacion no encontrada con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Ubicacion encontrada con ID:", id)
	RespondJSON(w, http.StatusOK, ubicacion)
}

func (s *Servidor) ActualizarUbicacion(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var actualizada models.Ubicacion
	err = json.NewDecoder(r.Body).Decode(&actualizada)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	ubicacion, err := s.UbicacionService.Actualizar(id, actualizada)
	if err != nil {
		if errors.Is(err, service.ErrProvinciaUbicacionObligatoria) || errors.Is(err, service.ErrCiudadUbicacionObligatoria) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, service.ErrUbicacionNoEncontrada) {
			fmt.Println("Ubicacion no encontrada con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Ubicacion actualizada con ID:", id)
	RespondJSON(w, http.StatusOK, ubicacion)
}

func (s *Servidor) EliminarUbicacion(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	err = s.UbicacionService.Eliminar(id)
	if err != nil {
		if errors.Is(err, service.ErrUbicacionNoEncontrada) {
			fmt.Println("Ubicacion no encontrada con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Ubicacion eliminada con ID:", id)
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Ubicacion eliminada correctamente"})
}
