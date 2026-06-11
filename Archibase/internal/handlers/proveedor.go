package handlers

import (
	"encoding/json"
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoProveedor)
}

func ObtenerProveedores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
			json.NewEncoder(w).Encode(proveedor)
			return
		}
	}

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
			json.NewEncoder(w).Encode(actualizado)
			return
		}
	}
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
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
