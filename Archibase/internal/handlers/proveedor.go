package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoProveedor models.Proveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoProveedor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevoProveedor.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	proveedorCreado := s.Almacen.CrearProveedor(nuevoProveedor)
	fmt.Println("--> Proveedor creado con ID:", proveedorCreado.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(proveedorCreado)
}

func (s *Servidor) ObtenerProveedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los proveedores")
	proveedores := s.Almacen.ListarProveedores()
	json.NewEncoder(w).Encode(proveedores)
}

func (s *Servidor) ObtenerProveedorPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	proveedor, encontrado := s.Almacen.BuscarProveedorPorID(id)
	if !encontrado {
		fmt.Println("--> Proveedor no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Proveedor no encontrado"})
		return
	}
	fmt.Println("--> Proveedor encontrado con ID:", id)
	json.NewEncoder(w).Encode(proveedor)
}

func (s *Servidor) ActualizarProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	var actualizado models.Proveedor
	err = json.NewDecoder(r.Body).Decode(&actualizado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if actualizado.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	proveedor, encontrado := s.Almacen.ActualizarProveedor(id, actualizado)
	if !encontrado {
		fmt.Println("--> Proveedor no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Proveedor no encontrado"})
		return
	}
	fmt.Println("--> Proveedor actualizado con ID:", id)
	json.NewEncoder(w).Encode(proveedor)
}

func (s *Servidor) EliminarProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	exito := s.Almacen.EliminarProveedor(id)
	if !exito {
		fmt.Println("--> Proveedor no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Proveedor no encontrado"})
		return
	}
	fmt.Println("--> Proveedor eliminado con ID:", id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Proveedor eliminado correctamente"})
}
