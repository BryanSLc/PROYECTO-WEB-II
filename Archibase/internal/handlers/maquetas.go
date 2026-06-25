package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

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

	maquetaCreada, svcErr := s.MaquetaSvc.Crear(nuevaMaqueta)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Maqueta creada con ID:", maquetaCreada.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(maquetaCreada)
}

func (s *Servidor) ObtenerMaquetas(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todas las maquetas")
	json.NewEncoder(respuesta).Encode(s.MaquetaSvc.Listar())
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

	maqueta, svcErr := s.MaquetaSvc.BuscarPorID(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
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

	maqueta, svcErr := s.MaquetaSvc.Actualizar(id, maquetaActualizada)
	if svcErr != nil {
		responderError(respuesta, svcErr)
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

	svcErr := s.MaquetaSvc.Eliminar(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
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

	evolucionCreada, svcErr := s.MaquetaSvc.AgregarEvolucion(nuevaEvolucion)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

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

	historial, svcErr := s.MaquetaSvc.ListarEvolucion(maquetaID)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Obteniendo historial de evolucion para maqueta:", maquetaID)
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

	svcErr := s.MaquetaSvc.EliminarEvolucion(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Evolucion eliminada con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Evolucion eliminada correctamente"})
}
