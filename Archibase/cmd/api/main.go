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
	almacen := storage.NuevoSQLiteStorage("./archibase.db")

	// 2. Inyectamos el almacén persistente al servidor de handlers
	servidor := handlers.NuevoServidor(almacen)
	enrutador := chi.NewRouter()

	// ====================================================
	// CONFIGURACIÓN GLOBAL DE INTERCEPTORES (MIDDLEWARES)
	// ====================================================
	// Habilitamos CORS de manera global para permitir peticiones desde navegadores o herramientas externas
	enrutador.Use(middleware.Cors)

	// ====================================================
	// ZONA 1: RUTAS PÚBLICAS (Sin Middleware / Acceso Libre)
	// ====================================================

	// Un estudiante nuevo DEBE poder registrarse sin un Token previo
	enrutador.Post("/api/v1/usuarios", servidor.CrearUsuario)

	// Aquí irá tu endpoint público de login en el futuro:
	// enrutador.Post("/api/v1/auth/login", servidor.Login)

	// ====================================================
	// ZONA 2: RUTAS PROTEGIDAS (Bajo la vigilancia del Middleware)
	// ====================================================
	enrutador.Group(func(r chi.Router) {
		// Activamos tu middleware de autenticación en español para proteger la lógica de negocio
		// Nota: En un paso posterior inyectaremos aquí tu servicio de autenticación (ej: servidor.AuthService)
		r.Use(middleware.AuthMiddleware)

		// --- MÓDULO USUARIOS (ADMINISTRACIÓN PROTEGIDA) ---
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

		// --- MÓDULOS CON FUNCIONES GLOBALES (handlers.*) ---

		r.Route("/api/v1/proveedores", func(r chi.Router) {
			r.Post("/", handlers.CrearProveedor)
			r.Get("/", handlers.ObtenerProveedores)
			r.Get("/{id}", handlers.ObtenerProveedorPorID)
			r.Put("/{id}", handlers.ActualizarProveedor)
			r.Delete("/{id}", handlers.EliminarProveedor)
		})

		r.Route("/api/v1/asesores", func(r chi.Router) {
			r.Get("/", servidor.ObtenerAsesores)
			r.Post("/", servidor.CrearAsesor)
			r.Get("/{id}", servidor.ObtenerAsesorPorID)
			r.Put("/{id}", servidor.ActualizarAsesor)
			r.Delete("/{id}", servidor.EliminarAsesor)
		})

		r.Route("/api/v1/contrataciones", func(r chi.Router) {
			r.Get("/", servidor.ObtenerContrataciones)
			r.Post("/", servidor.CrearContratacion)
			r.Get("/{id}", servidor.ObtenerContratacionPorID)
			r.Put("/{id}", servidor.ActualizarContratacion)
			r.Delete("/{id}", servidor.EliminarContratacion)
		})

		r.Route("/api/v1/ubicaciones", func(r chi.Router) {
			r.Post("/", handlers.CrearUbicacion)
			r.Get("/", handlers.ObtenerUbicaciones)
			r.Get("/{id}", handlers.ObtenerUbicacionPorID)
			r.Put("/", handlers.ActualizarUbicacion)
			r.Delete("/{id}", handlers.EliminarUbicacion)
		})

		r.Route("/api/v1/servicios", func(r chi.Router) {
			r.Get("/", servidor.ObtenerServicios)
			r.Post("/", servidor.CrearServicio)
			r.Get("/{id}", servidor.ObtenerServicioPorID)
			r.Put("/{id}", servidor.ActualizarServicio)
			r.Delete("/{id}", servidor.EliminarServicio)
		})

		r.Route("/api/v1/materiales", func(r chi.Router) {
			r.Post("/", handlers.CrearMaterial)
			r.Get("/", handlers.ObtenerMateriales)
			r.Get("/{id}", handlers.ObtenerMaterialPorID)
			r.Put("/{id}", handlers.ActualizarMaterial)
			r.Delete("/{id}", handlers.EliminarMaterial)
		})
	})

	fmt.Println("Servidor de Archibase corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", enrutador); err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
