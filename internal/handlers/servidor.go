package handlers

import (
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// Servidor ahora usa la interfaz storage.Almacen en vez de *storage.SQLiteStorage.
// Esto permite que tanto SQLiteStorage (tests) como PostgresStorage (Docker)
// se inyecten sin cambiar nada más.
type Servidor struct {
	Almacen          storage.Almacen
	AsesorService    *service.AsesorService
	ProveedorService *service.ProveedorService
	MaterialService  *service.MaterialService
	UsuarioService   *service.UsuarioService
	MaquetaService   *service.MaquetaService
	RecetaService    *service.RecetaService
	UbicacionService *service.UbicacionService
	AuthService      *service.AuthService
}

// NuevoServidor acepta cualquier implementación de storage.Almacen
func NuevoServidor(a storage.Almacen) *Servidor {
	return &Servidor{
		Almacen:          a,
		AsesorService:    service.NuevoAsesorService(a),
		ProveedorService: service.NuevoProveedorService(a),
		MaterialService:  service.NuevoMaterialService(a),
		UsuarioService:   service.NuevoUsuarioService(a),
		MaquetaService:   service.NuevoMaquetaService(a),
		RecetaService:    service.NuevoRecetaService(a),
		UbicacionService: service.NuevoUbicacionService(a),
		AuthService:      service.NuevoAuthService(a),
	}
}
