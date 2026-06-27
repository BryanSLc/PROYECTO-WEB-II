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

func (s *Servidor) CrearProveedor(w http.ResponseWriter, r *http.Request) {
	var nuevoProveedor models.Proveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoProveedor)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	proveedorCreado, err := s.ProveedorService.Crear(nuevoProveedor)
	if err != nil {
		if errors.Is(err, service.ErrNombreProveedorObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Proveedor creado con ID:", proveedorCreado.ID)
	RespondJSON(w, http.StatusCreated, proveedorCreado)
}

func (s *Servidor) ObtenerProveedores(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Obteniendo todos los proveedores")
	RespondJSON(w, http.StatusOK, s.ProveedorService.Listar())
}

func (s *Servidor) ObtenerProveedorPorID(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	proveedor, err := s.ProveedorService.BuscarPorID(id)
	if err != nil {
		if errors.Is(err, service.ErrProveedorNoEncontrado) {
			fmt.Println("Proveedor no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Proveedor encontrado con ID:", id)
	RespondJSON(w, http.StatusOK, proveedor)
}

func (s *Servidor) ActualizarProveedor(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var actualizado models.Proveedor
	err = json.NewDecoder(r.Body).Decode(&actualizado)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Datos invalidos")
		return
	}
	proveedor, err := s.ProveedorService.Actualizar(id, actualizado)
	if err != nil {
		if errors.Is(err, service.ErrNombreProveedorObligatorio) {
			RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, service.ErrProveedorNoEncontrado) {
			fmt.Println("Proveedor no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Proveedor actualizado con ID:", id)
	RespondJSON(w, http.StatusOK, proveedor)
}

func (s *Servidor) EliminarProveedor(w http.ResponseWriter, r *http.Request) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	err = s.ProveedorService.Eliminar(id)
	if err != nil {
		if errors.Is(err, service.ErrProveedorNoEncontrado) {
			fmt.Println("Proveedor no encontrado con ID:", id)
			RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondError(w, http.StatusInternalServerError, "Error interno del servidor")
		return
	}
	fmt.Println("Proveedor eliminado con ID:", id)
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "Proveedor eliminado correctamente"})
}
