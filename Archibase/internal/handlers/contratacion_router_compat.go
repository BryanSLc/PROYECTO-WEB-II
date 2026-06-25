package handlers

import "net/http"

// Compatibilidad: el router en cmd/api/main.go espera handlers.* para contrataciones.

func ObtenerContrataciones(w http.ResponseWriter, r *http.Request) {
	// Usar el tipo existente a través de un singleton no aplica; por ahora dejamos explícito que
	// el router está desalineado. Para compilar y correr el resto del proyecto, se mantiene como stub.
	w.WriteHeader(http.StatusNotImplemented)
}

func CrearContratacion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func ObtenerContratacionPorID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func ActualizarContratacion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func EliminarContratacion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
