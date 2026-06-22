package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

func CrearProveedor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoProveedor models.Proveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoProveedor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if nuevoProveedor.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nuevoProveedor.ID = storage.ConteoProveedores
	storage.ConteoProveedores++

	storage.ListaProveedores =
		append(storage.ListaProveedores, nuevoProveedor)
	fmt.Println("--> Proveedor creado con ID:", nuevoProveedor.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoProveedor)
}

func ObtenerProveedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los proveedores")
	json.NewEncoder(w).Encode(storage.ListaProveedores)
}

func ObtenerProveedorPorID(
	w http.ResponseWriter,
	r *http.Request,
) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, proveedor := range storage.ListaProveedores {
		if proveedor.ID == id {
			fmt.Println("--> Proveedor encontrado con ID:", id)
			json.NewEncoder(w).Encode(proveedor)
			return
		}
	}

	fmt.Println("--> Proveedor no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}

func ActualizarProveedor(
	w http.ResponseWriter,
	r *http.Request,
) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var actualizado models.Proveedor
	err = json.NewDecoder(r.Body).Decode(&actualizado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, proveedor := range storage.ListaProveedores {
		if proveedor.ID == id {
			actualizado.ID = id
			storage.ListaProveedores[i] = actualizado
			fmt.Println("--> Proveedor actualizado con ID:", id)
			json.NewEncoder(w).Encode(actualizado)
			return
		}
	}
	fmt.Println("--> Proveedor no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}

func EliminarProveedor(
	w http.ResponseWriter,
	r *http.Request,
) {
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, proveedor := range storage.ListaProveedores {
		if proveedor.ID == id {
			storage.ListaProveedores =
				append(
					storage.ListaProveedores[:i],
					storage.ListaProveedores[i+1:]...,
				)
			fmt.Println("--> Proveedor eliminado con ID:", id)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	fmt.Println("--> Proveedor no encontrado con ID:", id)
	w.WriteHeader(http.StatusNotFound)
}
