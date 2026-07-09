package storage

import "proyecto/internal/models"

// Almacen es la interfaz que tanto SQLiteStorage como PostgresStorage implementan.
// Los servicios y el Servidor usan esta interfaz, así los tests siguen
// funcionando con SQLiteStorage y producción usa PostgresStorage.
type Almacen interface {
	// Usuarios
	CrearUsuario(models.Usuario) models.Usuario
	ListarUsuarios() []models.Usuario
	BuscarUsuarioPorID(int) (models.Usuario, bool)
	BuscarUsuarioPorEmail(string) (models.Usuario, bool)
	ActualizarUsuario(int, models.Usuario) (models.Usuario, bool)
	EliminarUsuario(int) bool

	// Maquetas
	CrearMaqueta(models.Maqueta) models.Maqueta
	ListarMaquetas() []models.Maqueta
	BuscarMaquetaPorID(int) (models.Maqueta, bool)
	ActualizarMaqueta(int, models.Maqueta) (models.Maqueta, bool)
	EliminarMaqueta(int) bool

	// Evolución Maqueta
	AgregarEvolucion(models.EvolucionMaqueta) models.EvolucionMaqueta
	ListarEvolucionPorMaqueta(int) []models.EvolucionMaqueta
	EliminarEvolucion(int) bool

	// Recetas
	CrearReceta(models.Receta) models.Receta
	ListarRecetas() []models.Receta
	ListarRecetasPorMaqueta(int) []models.Receta
	BuscarRecetaPorID(int) (models.Receta, bool)
	ActualizarReceta(int, models.Receta) (models.Receta, bool)
	EliminarReceta(int) bool

	// Proveedores
	CrearProveedor(models.Proveedor) models.Proveedor
	ListarProveedores() []models.Proveedor
	BuscarProveedorPorID(int) (models.Proveedor, bool)
	ActualizarProveedor(int, models.Proveedor) (models.Proveedor, bool)
	EliminarProveedor(int) bool

	// Materiales
	CrearMaterial(models.MaterialProveedor) models.MaterialProveedor
	ListarMateriales() []models.MaterialProveedor
	BuscarMaterialPorID(int) (models.MaterialProveedor, bool)
	ActualizarMaterial(int, models.MaterialProveedor) (models.MaterialProveedor, bool)
	EliminarMaterial(int) bool

	// Ubicaciones
	CrearUbicacion(models.Ubicacion) models.Ubicacion
	ListarUbicaciones() []models.Ubicacion
	BuscarUbicacionPorID(int) (models.Ubicacion, bool)
	ActualizarUbicacion(int, models.Ubicacion) (models.Ubicacion, bool)
	EliminarUbicacion(int) bool

	// Asesores
	CrearAsesor(models.Asesor) models.Asesor
	ListarAsesores() []models.Asesor
	BuscarAsesorPorID(int) (models.Asesor, bool)
	ActualizarAsesor(int, models.Asesor) (models.Asesor, bool)
	EliminarAsesor(int) bool

	// Contrataciones
	CrearContratacion(models.Contratacion) models.Contratacion
	ListarContrataciones() []models.Contratacion
	BuscarContratacionPorID(int) (models.Contratacion, bool)
	ActualizarContratacion(int, models.Contratacion) (models.Contratacion, bool)
	EliminarContratacion(int) bool

	// Servicios
	CrearServicio(models.Servicio) models.Servicio
	ListarServicios() []models.Servicio
	BuscarServicioPorID(int) (models.Servicio, bool)
	ActualizarServicio(int, models.Servicio) (models.Servicio, bool)
	EliminarServicio(int) bool
}
