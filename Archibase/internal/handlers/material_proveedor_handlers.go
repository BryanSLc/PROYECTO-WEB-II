package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nuevoMaterial models.MaterialProveedor
	err := json.NewDecoder(r.Body).Decode(&nuevoMaterial)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevoMaterial.Nombre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	materialCreado := s.Almacen.CrearMaterial(nuevoMaterial)
	fmt.Println("--> Material creado con ID:", materialCreado.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(materialCreado)
}

func (s *Servidor) ObtenerMateriales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los materiales")
	materiales := s.Almacen.ListarMateriales()
	json.NewEncoder(w).Encode(materiales)
}

func (s *Servidor) ObtenerMaterialPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	material, encontrado := s.Almacen.BuscarMaterialPorID(id)
	if !encontrado {
		fmt.Println("--> Material no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Material no encontrado"})
		return
	}

	fmt.Println("--> Material encontrado con ID:", id)
	json.NewEncoder(w).Encode(material)
}

func (s *Servidor) ActualizarMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var actualizado models.MaterialProveedor
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

	material, encontrado := s.Almacen.ActualizarMaterial(id, actualizado)
	if !encontrado {
		fmt.Println("--> Material no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Material no encontrado"})
		return
	}

	fmt.Println("--> Material actualizado con ID:", id)
	json.NewEncoder(w).Encode(material)
}

func (s *Servidor) EliminarMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	exito := s.Almacen.EliminarMaterial(id)
	if !exito {
		fmt.Println("--> Material no encontrado con ID:", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Material no encontrado"})
		return
	}

	fmt.Println("--> Material eliminado con ID:", id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Material eliminado correctamente"})
}
