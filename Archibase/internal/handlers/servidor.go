package handlers

import (
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

// Servidor ahora agrupa las capas de servicios de negocio inyectadas
type Servidor struct {
	Almacen        *storage.SQLiteStorage
	UsuarioService *service.UsuarioService
	MaquetaService *service.MaquetaService
	RecetaService  *service.RecetaService
	AuthService    *service.AuthService
}

// NuevoServidor construye la instancia inyectando las dependencias requeridas
func NuevoServidor(a *storage.SQLiteStorage) *Servidor {
	return &Servidor{
		Almacen:        a,
		UsuarioService: service.NuevoUsuarioService(a),
		MaquetaService: service.NuevoMaquetaService(a),
		RecetaService:  service.NuevoRecetaService(a),
		AuthService:    service.NuevoAuthService(a),
	}
}
