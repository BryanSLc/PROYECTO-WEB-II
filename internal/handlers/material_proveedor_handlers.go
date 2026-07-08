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

func (s *Servidor) CrearMaterial(w http.ResponseWriter, r *http.Request) {
	var nuevoMaterial models.MaterialProveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoMaterial)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	materialCreado, err := s.MaterialService.Crear(nuevoMaterial)
	if err != nil {
		if errors.Is(err, service.ErrNombreMaterialObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Material creado con ID:", materialCreado.ID)
	RespondJSON(w, http.StatusCreated, materialCreado)
}

func (s *Servidor) ObtenerMateriales(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Obteniendo todos los materiales")
	RespondJSON(w, http.StatusOK, s.MaterialService.Listar())
}

func (s *Servidor) ObtenerMaterialPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	material, err := s.MaterialService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrMaterialNoEncontrado) {
			fmt.Println("Material no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Material encontrado con ID:", id)
	RespondJSON(w, http.StatusOK, material)
}

func (s *Servidor) ActualizarMaterial(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var actualizado models.MaterialProveedor
	err = json.NewDecoder(r.Body).Decode(&actualizado)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	material, err := s.MaterialService.Actualizar(id, actualizado)
	if err != nil {
		if errors.Is(err, service.ErrNombreMaterialObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, service.ErrMaterialNoEncontrado) {
			fmt.Println("Material no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Material actualizado con ID:", id)
	RespondJSON(w, http.StatusOK, material)
}

func (s *Servidor) EliminarMaterial(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	err = s.MaterialService.Eliminar(id)
	if err != nil {
		if errors.Is(err, service.ErrMaterialNoEncontrado) {
			fmt.Println("Material no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Material eliminado con ID:", id)
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Material eliminado correctamente"})
}
