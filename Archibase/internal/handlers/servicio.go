package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "proyecto/internal/models"
    "proyecto/internal/storage"
)

func GetAllServicios(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(storage.GetAllServicios())
}

func GetServicioByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
    }
    servicio, err := storage.GetServicioByID(id)
    if err != nil {
        http.Error(w, `{"error":"servicio no encontrado"}`, http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(servicio)
}

func CreateServicio(w http.ResponseWriter, r *http.Request) {
    var s models.Servicio
    if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
        http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
        return
    }
    if s.Titulo == "" || s.Descripcion == "" || s.Disponibilidad == "" || s.IDasesor == 0 {
        http.Error(w, `{"error":"titulo, descripcion, disponibilidad e id_asesor son requeridos"}`, http.StatusBadRequest)
        return
    }
    if s.Precio < 0 {
        http.Error(w, `{"error":"el precio no puede ser negativo"}`, http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(storage.CreateServicio(s))
}

func UpdateServicio(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
    }
    var s models.Servicio
    if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
        http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
        return
    }
    if s.Titulo == "" || s.Descripcion == "" || s.Disponibilidad == "" || s.IDasesor == 0 {
        http.Error(w, `{"error":"titulo, descripcion, disponibilidad e id_asesor son requeridos"}`, http.StatusBadRequest)
        return
    }
    actualizado, err := storage.UpdateServicio(id, s)
    if err != nil {
        http.Error(w, `{"error":"servicio no encontrado"}`, http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(actualizado)
}

func DeleteServicio(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
   