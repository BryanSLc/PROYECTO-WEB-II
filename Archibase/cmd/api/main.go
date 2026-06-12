package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/handlers"
)

func main() {
	enrutador := chi.NewRouter()

	// Rutas de maquetas
	enrutador.Route("/api/v1/maquetas", func(r chi.Router) {
		r.Post("/", handlers.CrearMaqueta)
		r.Get("/", handlers.ObtenerMaquetas)
		r.Get("/{id}", handlers.ObtenerMaquetaPorID)
		r.Put("/{id}", handlers.ActualizarMaqueta)
		r.Delete("/{id}", handlers.EliminarMaqueta)
	})

	// Rutas de proveedores
	enrutador.Route("/api/v1/proveedores", func(r chi.Router) {
		r.Post("/", handlers.CrearProveedor)
		r.Get("/", handlers.ObtenerProveedores)
		r.Get("/{id}", handlers.ObtenerProveedorPorID)
		r.Put("/{id}", handlers.ActualizarProveedor)
		r.Delete("/{id}", handlers.EliminarProveedor)
	})

	// Rutas de asesores
	enrutador.Route("/api/v1/asesores", func(r chi.Router) {
		r.Post("/", handlers.CrearAsesor)
		r.Get("/", handlers.ObtenerAsesores)
		r.Get("/{id}", handlers.ObtenerAsesorPorID)
		r.Put("/{id}", handlers.ActualizarAsesor)
		r.Delete("/{id}", handlers.EliminarAsesor)
	})

	// Rutas de contrataciones
	enrutador.Route("/api/v1/contrataciones", func(r chi.Router) {
		r.Post("/", handlers.CreateContratacion)
		r.Get("/", handlers.GetAllContrataciones)
		r.Get("/{id}", handlers.GetContratacionByID)
		r.Put("/{id}", handlers.UpdateContratacion)
		r.Delete("/{id}", handlers.DeleteContratacion)
	})

	// Rutas de ubicaciones
	enrutador.Route("/api/v1/ubicaciones", func(r chi.Router) {
		r.Post("/", handlers.CrearUbicacion)
		r.Get("/", handlers.ObtenerUbicaciones)
		r.Get("/{id}", handlers.ObtenerUbicacionPorID)
		r.Put("/{id}", handlers.ActualizarUbicacion)
		r.Delete("/{id}", handlers.EliminarUbicacion)
	})

	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")

	if err := http.ListenAndServe(":8080", enrutador); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}