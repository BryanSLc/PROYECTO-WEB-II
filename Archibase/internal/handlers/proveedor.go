package handlers

import (
	"encoding/json"
	"net/http"

	"proyecto/internal/models"
	"proyecto/internal/storage"
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
