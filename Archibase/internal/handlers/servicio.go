package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"proyecto/internal/models"
	"proyecto/internal/storage"

	"github.com/go-chi/chi/v5"
)

/func CrearServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	var nuevoServicio models.Servicio
	lector := json.NewDecoder(peticion.Body)
	err := lector.Decode(&nuevoServicio)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if nuevoServicio.Nombre == "" { // Asumiendo 'Nombre' como campo obligatorio. Ajusta según tu struct.
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	nuevoServicio.ID = storage.ConteoServicios
	storage.ConteoServicios++
	storage.ListaServicios = append(storage.ListaServicios, nuevoServicio)
	fmt.Println("--> Servicio creado con ID:", nuevoServicio.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(nuevoServicio)
}

func ObtenerServicios(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los servicios")
	json.NewEncoder(respuesta).Encode(storage.ListaServicios)
}

func ObtenerServicioPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for _, servicio := range storage.ListaServicios {
		if servicio.ID == id {
			fmt.Println("--> Servicio encontrado con ID:", id)
			json.NewEncoder(respuesta).Encode(servicio)
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound) // error 404
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
}

func ActualizarServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	var servicioActualizado models.Servicio
	lector := json.NewDecoder(peticion.Body)
	err = lector.Decode(&servicioActualizado)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}
	if servicioActualizado.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest) // 400 error
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	for i, servicio := range storage.ListaServicios {
		if servicio.ID == id {
			servicioActualizado.ID = id
			storage.ListaServicios[i] = servicioActualizado
			fmt.Println("--> Servicio actualizado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]interface{}{
				"mensaje":  "Servicio actualizado",
				"servicio": servicioActualizado,
			})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
}

func EliminarServicio(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}
	for i, servicio := range storage.ListaServicios {
		if servicio.ID == id {
			storage.ListaServicios = append(storage.ListaServicios[:i], storage.ListaServicios[i+1:]...)

			fmt.Println("--> Servicio eliminado con ID:", id)
			json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Servicio eliminado correctamente"})
			return
		}
	}
	respuesta.WriteHeader(http.StatusNotFound)
	json.NewEncoder(respuesta).Encode(map[string]string{"error": "Servicio no encontrado"})
}