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
	})
	// Rutas de proveedor
	enrutador.Route("/api/v1/proveedores", func(r chi.Router) {
		r.Post("/", handlers.CrearProveedor)
		r.Get("/", handlers.ObtenerProveedores)
		r.Get("/{id}", handlers.ObtenerProveedorPorID)
		r.Put("/{id}", handlers.ActualizarProveedor)
		r.Delete("/{id}", handlers.EliminarProveedor)
	})
	// Rutas de asesor
	enrutador.Route("/api/v1/asesores", func(r chi.Router) {
		r.Get("/", handlers.GetAllAsesores)
		r.Post("/", handlers.CreateAsesor)
		r.Get("/{id}", handlers.GetAsesorByID)
		r.Put("/{id}", handlers.UpdateAsesor)
		r.Delete("/{id}", handlers.DeleteAsesor)
	})
	// Rutas de contratacion
	enrutador.Route("/api/v1/contrataciones", func(r chi.Router) {
		  r.Get("/", handlers.GetAllContrataciones)
            r.Post("/", handlers.CreateContratacion)
            r.Get("/{id}", handlers.GetContratacionByID)
            r.Put("/{id}", handlers.UpdateContratacion)
            r.Delete("/{id}", handlers.DeleteContratacion)
        })
	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", enrutador)
	
    }