package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/handlers"
	"proyecto/internal/handlers/middleware"
	"proyecto/internal/storage"
)

func main() {
	// 1. Inicializamos el almacenamiento en SQLite (creará automáticamente el archivo archibase.db)
	almacen := storage.NuevoStorage()

	// 2. Inyectamos el almacén persistente al servidor de handlers
	servidor := handlers.NuevoServidor(almacen)
	enrutador := chi.NewRouter()

	// ====================================================
	// CONFIGURACIÓN GLOBAL DE INTERCEPTORES (MIDDLEWARES)
	// ====================================================
	enrutador.Use(middleware.Cors)

	// ====================================================
	// ZONA 1: RUTAS PÚBLICAS (Sin Middleware / Acceso Libre)
	// ====================================================
	enrutador.Post("/api/v1/usuarios", servidor.CrearUsuario)

	enrutador.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/registro", servidor.Registrar)
		r.Post("/login", servidor.Login)
	})

	// ====================================================
	// ZONA 2: RUTAS PROTEGIDAS (Bajo la vigilancia del Middleware)
	// ====================================================
	enrutador.Group(func(r chi.Router) {
		// El middleware necesita el AuthService para verificar el token
		r.Use(middleware.AuthMiddleware(servidor.AuthService))

		// --- MÓDULO USUARIOS ---
		r.Route("/api/v1/usuarios", func(r chi.Router) {
			r.Get("/", servidor.ObtenerUsuarios)
			r.Get("/{id}", servidor.ObtenerUsuarioPorID)
			r.Put("/{id}", servidor.ActualizarUsuario)
			r.Delete("/{id}", servidor.EliminarUsuario)
		})

		// --- MÓDULO MAQUETAS ---
		r.Route("/api/v1/maquetas", func(r chi.Router) {
			r.Post("/", servidor.CrearMaqueta)
			r.Get("/", servidor.ObtenerMaquetas)
			r.Get("/{id}", servidor.ObtenerMaquetaPorID)
			r.Put("/{id}", servidor.ActualizarMaqueta)
			r.Delete("/{id}", servidor.EliminarMaqueta)

			r.Post("/evolucion", servidor.AgregarEvolucionMaqueta)
			r.Get("/{id}/evolucion", servidor.ObtenerEvolucionPorMaqueta)
			r.Delete("/evolucion/{id}", servidor.EliminarEvolucion)
		})

		// --- MÓDULO RECETAS ---
		r.Route("/api/v1/recetas", func(r chi.Router) {
			r.Post("/", servidor.CrearReceta)
			r.Get("/", servidor.ObtenerRecetas)
			r.Get("/{id}", servidor.ObtenerRecetaPorID)
			r.Put("/{id}", servidor.ActualizarReceta)
			r.Delete("/{id}", servidor.EliminarReceta)
		})

		// --- MÓDULO PROVEEDORES ---
		r.Route("/api/v1/proveedores", func(r chi.Router) {
			r.Post("/", servidor.CrearProveedor)
			r.Get("/", servidor.ObtenerProveedores)
			r.Get("/{id}", servidor.ObtenerProveedorPorID)
			r.Put("/{id}", servidor.ActualizarProveedor)
			r.Delete("/{id}", servidor.EliminarProveedor)
		})

		// --- MÓDULO UBICACIONES ---
		r.Route("/api/v1/ubicaciones", func(r chi.Router) {
			r.Post("/", servidor.CrearUbicacion)
			r.Get("/", servidor.ObtenerUbicaciones)
			r.Get("/{id}", servidor.ObtenerUbicacionPorID)
			r.Put("/{id}", servidor.ActualizarUbicacion)
			r.Delete("/{id}", servidor.EliminarUbicacion)
		})

		// --- MÓDULO MATERIALES ---
		r.Route("/api/v1/materiales", func(r chi.Router) {
			r.Post("/", servidor.CrearMaterial)
			r.Get("/", servidor.ObtenerMateriales)
			r.Get("/{id}", servidor.ObtenerMaterialPorID)
			r.Put("/{id}", servidor.ActualizarMaterial)
			r.Delete("/{id}", servidor.EliminarMaterial)
		})

		// --- MÓDULO ASESORES (una sola vez) ---
		r.Route("/api/v1/asesores", func(r chi.Router) {
			r.Get("/", handlers.GetAllAsesores)
			r.Post("/", handlers.CreateAsesor)
			r.Get("/{id}", handlers.GetAsesorByID)
			r.Put("/{id}", handlers.UpdateAsesor)
			r.Delete("/{id}", handlers.DeleteAsesor)
		})

		// --- MÓDULO SERVICIOS (una sola vez) ---
		r.Route("/api/v1/servicios", func(r chi.Router) {
			r.Get("/", handlers.ObtenerServicios)
			r.Post("/", handlers.CrearServicio)
			r.Get("/{id}", handlers.ObtenerServicioPorID)
			r.Put("/{id}", handlers.ActualizarServicio)
			r.Delete("/{id}", handlers.EliminarServicio)
		})

		// --- MÓDULO CONTRATACIONES (una sola vez) ---
		r.Route("/api/v1/contrataciones", func(r chi.Router) {
			r.Get("/", handlers.ObtenerContrataciones)
			r.Post("/", handlers.CrearContratacion)
			r.Get("/{id}", handlers.ObtenerContratacionPorID)
			r.Put("/{id}", handlers.ActualizarContratacion)
			r.Delete("/{id}", handlers.EliminarContratacion)
		})
	})

	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", enrutador); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
