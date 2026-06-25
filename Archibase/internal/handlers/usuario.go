package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/models"
)

func (s *Servidor) CrearUsuario(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	var nuevoUsuario models.Usuario
	lector := json.NewDecoder(peticion.Body)
	if err := lector.Decode(&nuevoUsuario); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if nuevoUsuario.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}
	if nuevoUsuario.Email == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El correo electronico es obligatorio"})
		return
	}

	usuarioCreado, svcErr := s.UsuarioSvc.Crear(nuevoUsuario)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Usuario creado con ID:", usuarioCreado.ID)
	respuesta.WriteHeader(http.StatusCreated)
	json.NewEncoder(respuesta).Encode(usuarioCreado)
}

func (s *Servidor) ObtenerUsuarios(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")
	fmt.Println("--> Obteniendo todos los usuarios")
	json.NewEncoder(respuesta).Encode(s.UsuarioSvc.Listar())
}

func (s *Servidor) ObtenerUsuarioPorID(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	usuario, svcErr := s.UsuarioSvc.BuscarPorID(id)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Usuario encontrado con ID:", id)
	json.NewEncoder(respuesta).Encode(usuario)
}

func (s *Servidor) ActualizarUsuario(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	var usuarioActualizado models.Usuario
	lector := json.NewDecoder(peticion.Body)
	if err := lector.Decode(&usuarioActualizado); err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "Datos invalidos"})
		return
	}

	if usuarioActualizado.Nombre == "" {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "El nombre es obligatorio"})
		return
	}

	usuario, svcErr := s.UsuarioSvc.Actualizar(id, usuarioActualizado)
	if svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Usuario actualizado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]interface{}{
		"mensaje": "Usuario actualizado correctamente",
		"usuario": usuario,
	})
}

func (s *Servidor) EliminarUsuario(respuesta http.ResponseWriter, peticion *http.Request) {
	respuesta.Header().Set("Content-Type", "application/json")

	idTexto := chi.URLParam(peticion, "id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		respuesta.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(respuesta).Encode(map[string]string{"error": "ID invalido"})
		return
	}

	if svcErr := s.UsuarioSvc.Eliminar(id); svcErr != nil {
		responderError(respuesta, svcErr)
		return
	}

	fmt.Println("--> Usuario eliminado con ID:", id)
	json.NewEncoder(respuesta).Encode(map[string]string{"mensaje": "Usuario eliminado correctamente"})
}
