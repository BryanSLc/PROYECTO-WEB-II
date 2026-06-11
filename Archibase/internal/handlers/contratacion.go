package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "proyecto/internal/models"
    "proyecto/internal/storage"
)

func GetAllContrataciones(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(storage.GetAllContrataciones())
}

func GetContratacionByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
    }
    contratacion, err := storage.GetContratacionByID(id)
    if err != nil {
        http.Error(w, `{"error":"contratacion no encontrada"}`, http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(contratacion)
}

func CreateContratacion(w http.ResponseWriter, r *http.Request) {
    var c models.Contratacion
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
        return
    }
    if c.Estudiante == "" || c.Fecha == "" || c.Estado == "" || c.IDservicio == 0 {
        http.Error(w, `{"error":"estudiante, fecha, estado e id_servicio son requeridos"}`, http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(storage.CreateContratacion(c))
}

func UpdateContratacion(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
    }
    var c models.Contratacion
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, `{"error":"body invalido"}`, http.StatusBadRequest)
        return
    }
    if c.Estudiante == "" || c.Fecha == "" || c.Estado == "" || c.IDservicio == 0 {
        http.Error(w, `{"error":"estudiante, fecha, estado e id_servicio son requeridos"}`, http.StatusBadRequest)
        return
    }
    actualizada, err := storage.UpdateContratacion(id, c)
    if err != nil {
        http.Error(w, `{"error":"contratacion no encontrada"}`, http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(actualizada)
}

func DeleteContratacion(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, `{"error":"ID invalido"}`, http.StatusBadRequest)
        return
    }
    if err := storage.DeleteContratacion(id); err != nil {
        http.Error(w, `{"error":"contratacion no encontrada"}`, http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"mensaje": "contratacion eliminada"})
}