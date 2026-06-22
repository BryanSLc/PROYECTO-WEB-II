package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
	"proyecto/internal/storage"
)

type Servidor struct {
	Almacen *storage.SQLiteStorage
}

func NuevoServidor(a *storage.SQLiteStorage) *Servidor {
	return &Servidor{Almacen: a}
}

func (s *Servidor) CrearMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevaMaqueta models.Maqueta
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaMaqueta)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaMaqueta.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}
	maquetaCreada := s.Almacen.CrearMaqueta(nuevaMaqueta)
	fmt.Println("--> Maqueta creada con ID:", maquetaCreada.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(maquetaCreada)
}

func (s *Servidor) ObtenerMaquetas(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las maquetas")
	json.NewEncoder(respuesta).Encode(s.Almacen.ListarMaquetas())
}

func (s *Servidor) ObtenerMaquetaPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	maqueta, encontrada := s.Almacen.BuscarMaquetaPorID(id)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
		return
	}

	fmt.Println("--> Maqueta encontrada con ID:", id)
	json.NewEncoder(respuesta).Encode(maqueta)
}

func (s *Servidor) ActualizarMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var maquetaActualizada models.Maqueta
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&maquetaActualizada)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if maquetaActualizada.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo es obligatorio"})
		return
	}

	maqueta, encontrada := s.Almacen.ActualizarMaqueta(id, maquetaActualizada)
	if !encontrada {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
		return
	}

	fmt.Println("--> Maqueta actualizada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje": "Maqueta actualizada",
		"maqueta": maqueta,
	})
}

func (s *Servidor) EliminarMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if !s.Almacen.EliminarMaqueta(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
		return
	}

	fmt.Println("--> Maqueta asociada eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Maqueta eliminada correctamente"})
}

func (s *Servidor) AgregarEvolucionMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevaEvolucion models.EvolucionMaqueta
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevaEvolucion)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevaEvolucion.MaquetaID == 0 {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El ID de la maqueta es obligatorio"})
		return
	}
	if nuevaEvolucion.Titulo == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El titulo del avance es obligatorio"})
		return
	}
	if nuevaEvolucion.Paso <= 0 {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El numero de paso debe ser mayor a 0"})
		return
	}
	_, existe := s.Almacen.BuscarMaquetaPorID(nuevaEvolucion.MaquetaID)
	if !existe {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "La maqueta especificada no existe"})
		return
	}
	evolucionCreada := s.Almacen.AgregarEvolucion(nuevaEvolucion)
	fmt.Println("--> Evolucion registrada para Maqueta ID:", evolucionCreada.MaquetaID, "Paso:", evolucionCreada.Paso)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(evolucionCreada)
}

func (s *Servidor) ObtenerEvolucionPorMaqueta(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	maquetaID, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID de maqueta invalido"})
		return
	}
	_, existe := s.Almacen.BuscarMaquetaPorID(maquetaID)
	if !existe {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Maqueta no encontrada"})
		return
	}
	fmt.Println("--> Obteniendo historial de evolucion para maqueta:", maquetaID)
	historial := s.Almacen.ListarEvolucionPorMaqueta(maquetaID)
	json.NewEncoder(respuesta).Encode(historial)
}

func (s *Servidor) EliminarEvolucion(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if !s.Almacen.EliminarEvolucion(id) {
		respuesta.WriteHeader(http.StatusNotFound)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Evolucion no encontrada"})
		return
	}

	fmt.Println("--> Evolucion eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Evolucion eliminada correctamente"})
}
