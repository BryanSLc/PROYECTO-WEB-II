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
		r.Get("/", handlers.GetAllAsesores)
		r.Post("/", handlers.CreateAsesor)
		r.Get("/{id}", handlers.GetAsesorByID)
		r.Put("/{id}", handlers.UpdateAsesor)
		r.Delete("/{id}", handlers.DeleteAsesor)
	})
	// Rutas de contrataciones
	enrutador.Route("/api/v1/contrataciones", func(r chi.Router) {
		r.Get("/", handlers.ObtenerContrataciones)
		r.Post("/", handlers.CrearContratacion)
		r.Get("/{id}", handlers.ObtenerContratacionPorID)
		r.Put("/{id}", handlers.ActualizarContratacion)
		r.Delete("/{id}", handlers.EliminarContratacion)
	})
	// Rutas de ubicacion
	enrutador.Route("/api/v1/ubicaciones", func(r chi.Router) {
		r.Post("/", handlers.CrearUbicacion)
		r.Get("/", handlers.ObtenerUbicaciones)
		r.Get("/{id}", handlers.ObtenerUbicacionPorID)
		r.Put("/{id}", handlers.ActualizarUbicacion)
		r.Delete("/{id}", handlers.EliminarUbicacion)
	})
	// Rutas de servicios
	enrutador.Route("/api/v1/servicios", func(r chi.Router) {
		r.Get("/", handlers.ObtenerServicios)
		r.Post("/", handlers.CrearServicio)
		r.Get("/{id}", handlers.ObtenerServicioPorID)
		r.Put("/{id}", handlers.ActualizarServicio)
		r.Delete("/{id}", handlers.EliminarServicio)
	})
	// Rutas de material_proveedor
	enrutador.Route("/api/v1/materiales", func(r chi.Router) {
		r.Post("/", handlers.CrearMaterial)
		r.Get("/", handlers.ObtenerMateriales)
		r.Get("/{id}", handlers.ObtenerMaterialPorID)
		r.Put("/{id}", handlers.ActualizarMaterial)
		r.Delete("/{id}", handlers.EliminarMaterial)
	})
	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", enrutador)

	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")

	if err := http.ListenAndServe(":8080", enrutador); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
