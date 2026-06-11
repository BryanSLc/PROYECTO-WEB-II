package main

import (
	"fmt"

	"net/http"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/handlers"
)

func main() {
	enrutador := chi.NewRouter()
	enrutador.Route("/api/v1/maquetas", func(r chi.Router) {
		r.Post("/", handlers.CrearMaqueta)
		r.Get("/", handlers.ObtenerMaquetas)
		r.Get("/{id}", handlers.ObtenerMaquetaPorID)
		r.Put("/{id}", handlers.ActualizarMaqueta)
		// Rutas de proveedor
		r.Post("/", handlers.CrearProveedor)
		r.Get("/", handlers.ObtenerProveedores)
		r.Get("/{id}", handlers.ObtenerProveedorPorID)
		r.Put("/{id}", handlers.ActualizarProveedor)
		r.Delete("/{id}", handlers.EliminarProveedor)
	})
	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", enrutador)
}
