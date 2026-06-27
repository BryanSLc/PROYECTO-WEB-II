package handlers

import (
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// Servidor ahora agrupa las capas de servicios de negocio inyectadas
type Servidor struct {
	Almacen          *storage.SQLiteStorage
	ProveedorService *service.ProveedorService
	MaterialService  *service.MaterialService
	UsuarioService   *service.UsuarioService
	MaquetaService   *service.MaquetaService
	RecetaService    *service.RecetaService
	UbicacionService *service.UbicacionService
}

// NuevoServidor construye la instancia inyectando las dependencias requeridas
func NuevoServidor(a *storage.SQLiteStorage) *Servidor {
	return &Servidor{
		Almacen:          a,
		ProveedorService: service.NuevoProveedorService(a),
		MaterialService:  service.NuevoMaterialService(a),
		UsuarioService:   service.NuevoUsuarioService(a),
		MaquetaService:   service.NuevoMaquetaService(a),
		RecetaService:    service.NuevoRecetaService(a),
		UbicacionService: service.NuevoUbicacionService(a),
	}
}
