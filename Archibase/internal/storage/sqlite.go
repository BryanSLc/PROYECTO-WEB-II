package storage

import (
	"log"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"proyecto/internal/models"
)

type SQLiteStorage struct {
	db *gorm.DB
}

// NuevoSQLiteStorage inicializa la base de datos, crea las tablas y aplica parches necesarios
func NuevoSQLiteStorage(pathDB string) *SQLiteStorage {
	db, err := gorm.Open(sqlite.Open(pathDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a SQLite: %v", err)
	}

	// Migración automática de todos tus modelos
	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Maqueta{},
		&models.EvolucionMaqueta{},
		&models.Receta{},
		&models.Proveedor{},
		&models.MaterialProveedor{},
	)
	if err != nil {
		log.Fatalf("Error al realizar la migración: %v", err)
	}

	return &SQLiteStorage{db: db}
}

// ==========================================
//          MÉTODOS PARA PROVEEDORES
// ==========================================

func (s *SQLiteStorage) CrearProveedor(proveedor models.Proveedor) models.Proveedor {
	s.db.Create(&proveedor)
	return proveedor
}

func (s *SQLiteStorage) ListarProveedores() []models.Proveedor {
	var proveedores []models.Proveedor
	s.db.Find(&proveedores)
	return proveedores
}

func (s *SQLiteStorage) BuscarProveedorPorID(id int) (models.Proveedor, bool) {
	var proveedor models.Proveedor
	err := s.db.First(&proveedor, id).Error
	if err != nil {
		return models.Proveedor{}, false
	}
	return proveedor, true
}

func (s *SQLiteStorage) ActualizarProveedor(id int, datos models.Proveedor) (models.Proveedor, bool) {
	var proveedor models.Proveedor
	err := s.db.First(&proveedor, id).Error
	if err != nil {
		return models.Proveedor{}, false
	}

	proveedor.Nombre = datos.Nombre
	proveedor.Ciudad = datos.Ciudad
	proveedor.Provincia = datos.Provincia
	proveedor.Direccion = datos.Direccion
	proveedor.Telefono = datos.Telefono

	s.db.Save(&proveedor)
	return proveedor, true
}

func (s *SQLiteStorage) EliminarProveedor(id int) bool {
	resultado := s.db.Delete(&models.Proveedor{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//     MÉTODOS PARA MATERIAL PROVEEDOR
// ==========================================

func (s *SQLiteStorage) CrearMaterial(material models.MaterialProveedor) models.MaterialProveedor {
	s.db.Create(&material)
	return material
}

func (s *SQLiteStorage) ListarMateriales() []models.MaterialProveedor {
	var materiales []models.MaterialProveedor
	s.db.Find(&materiales)
	return materiales
}

func (s *SQLiteStorage) BuscarMaterialPorID(id int) (models.MaterialProveedor, bool) {
	var material models.MaterialProveedor
	err := s.db.First(&material, id).Error
	if err != nil {
		return models.MaterialProveedor{}, false
	}
	return material, true
}

func (s *SQLiteStorage) ActualizarMaterial(id int, datos models.MaterialProveedor) (models.MaterialProveedor, bool) {
	var material models.MaterialProveedor
	err := s.db.First(&material, id).Error
	if err != nil {
		return models.MaterialProveedor{}, false
	}

	material.Nombre = datos.Nombre
	material.Categoria = datos.Categoria
	material.PrecioReferencial = datos.PrecioReferencial
	material.IDProveedor = datos.IDProveedor

	s.db.Save(&material)
	return material, true
}

func (s *SQLiteStorage) EliminarMaterial(id int) bool {
	resultado := s.db.Delete(&models.MaterialProveedor{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA UBICACIONES
// ==========================================

func (s *SQLiteStorage) CrearUbicacion(ubicacion models.Ubicacion) models.Ubicacion {
	s.db.Create(&ubicacion)
	return ubicacion
}

func (s *SQLiteStorage) ListarUbicaciones() []models.Ubicacion {
	var ubicaciones []models.Ubicacion
	s.db.Find(&ubicaciones)
	return ubicaciones
}

func (s *SQLiteStorage) BuscarUbicacionPorID(id int) (models.Ubicacion, bool) {
	var ubicacion models.Ubicacion
	err := s.db.First(&ubicacion, id).Error
	if err != nil {
		return models.Ubicacion{}, false
	}
	return ubicacion, true
}

func (s *SQLiteStorage) ActualizarUbicacion(id int, datos models.Ubicacion) (models.Ubicacion, bool) {
	var ubicacion models.Ubicacion
	err := s.db.First(&ubicacion, id).Error
	if err != nil {
		return models.Ubicacion{}, false
	}

	ubicacion.Provincia = datos.Provincia
	ubicacion.Ciudad = datos.Ciudad

	s.db.Save(&ubicacion)
	return ubicacion, true
}

func (s *SQLiteStorage) EliminarUbicacion(id int) bool {
	resultado := s.db.Delete(&models.Ubicacion{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA MAQUETAS
// ==========================================

func (s *SQLiteStorage) CrearMaqueta(maqueta models.Maqueta) models.Maqueta {
	s.db.Create(&maqueta)
	return maqueta
}

func (s *SQLiteStorage) ListarMaquetas() []models.Maqueta {
	var maquetas []models.Maqueta
	s.db.Find(&maquetas)
	return maquetas
}

func (s *SQLiteStorage) BuscarMaquetaPorID(id int) (models.Maqueta, bool) {
	var maqueta models.Maqueta
	err := s.db.First(&maqueta, id).Error
	if err != nil {
		return models.Maqueta{}, false
	}
	return maqueta, true
}

func (s *SQLiteStorage) ActualizarMaqueta(id int, datos models.Maqueta) (models.Maqueta, bool) {
	var maqueta models.Maqueta
	err := s.db.First(&maqueta, id).Error
	if err != nil {
		return models.Maqueta{}, false
	}

	// Actualizamos los campos individuales
	maqueta.UsuarioID = datos.UsuarioID
	maqueta.Titulo = datos.Titulo
	maqueta.Descripcion = datos.Descripcion
	maqueta.Escala = datos.Escala
	maqueta.Materiales = datos.Materiales
	maqueta.Dimensiones = datos.Dimensiones

	s.db.Save(&maqueta)
	return maqueta, true
}

func (s *SQLiteStorage) EliminarMaqueta(id int) bool {
	resultado := s.db.Delete(&models.Maqueta{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//       MÉTODOS PARA EVOLUCIÓN MAQUETA
// ==========================================

func (s *SQLiteStorage) AgregarEvolucion(evolucion models.EvolucionMaqueta) models.EvolucionMaqueta {
	evolucion.Fecha = time.Now()
	s.db.Create(&evolucion)
	return evolucion
}

func (s *SQLiteStorage) ListarEvolucionPorMaqueta(maquetaID int) []models.EvolucionMaqueta {
	var historial []models.EvolucionMaqueta
	s.db.Where("maqueta_id = ?", maquetaID).Find(&historial)
	return historial
}

func (s *SQLiteStorage) EliminarEvolucion(id int) bool {
	resultado := s.db.Delete(&models.EvolucionMaqueta{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA USUARIOS
// ==========================================

func (s *SQLiteStorage) CrearUsuario(usuario models.Usuario) models.Usuario {
	usuario.FechaCreacion = time.Now()
	s.db.Create(&usuario)
	return usuario
}

func (s *SQLiteStorage) ListarUsuarios() []models.Usuario {
	var usuarios []models.Usuario
	s.db.Find(&usuarios)
	return usuarios
}

func (s *SQLiteStorage) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var usuario models.Usuario
	err := s.db.First(&usuario, id).Error
	if err != nil {
		return models.Usuario{}, false
	}
	return usuario, true
}

func (s *SQLiteStorage) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	var usuario models.Usuario
	err := s.db.First(&usuario, id).Error
	if err != nil {
		return models.Usuario{}, false
	}
	usuario.Nombre = datos.Nombre
	usuario.Apellido = datos.Apellido
	usuario.Email = datos.Email
	if datos.Password != "" {
		usuario.Password = datos.Password
	}
	usuario.Semestre = datos.Semestre
	usuario.Telefono = datos.Telefono
	usuario.Rol = datos.Rol

	s.db.Save(&usuario)
	return usuario, true
}

func (s *SQLiteStorage) EliminarUsuario(id int) bool {
	resultado := s.db.Delete(&models.Usuario{}, id)
	return resultado.RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA RECETAS
// ==========================================

func (s *SQLiteStorage) CrearReceta(receta models.Receta) models.Receta {
	s.db.Create(&receta)
	return receta
}

func (s *SQLiteStorage) ListarRecetas() []models.Receta {
	var recetas []models.Receta
	s.db.Find(&recetas)
	return recetas
}

func (s *SQLiteStorage) ListarRecetasPorMaqueta(maquetaID int) []models.Receta {
	var recetas []models.Receta
	s.db.Where("maqueta_id = ?", maquetaID).Find(&recetas)
	return recetas
}

func (s *SQLiteStorage) BuscarRecetaPorID(id int) (models.Receta, bool) {
	var receta models.Receta
	err := s.db.First(&receta, id).Error
	if err != nil {
		return models.Receta{}, false
	}
	return receta, true
}

func (s *SQLiteStorage) ActualizarReceta(id int, datos models.Receta) (models.Receta, bool) {
	var receta models.Receta
	err := s.db.First(&receta, id).Error
	if err != nil {
		return models.Receta{}, false
	}

	receta.MaquetaID = datos.MaquetaID
	receta.Titulo = datos.Titulo
	receta.Descripcion = datos.Descripcion
	receta.Pasos = datos.Pasos // GORM lo serializará automáticamente como JSON texto

	s.db.Save(&receta)
	return receta, true
}

func (s *SQLiteStorage) EliminarReceta(id int) bool {
	resultado := s.db.Delete(&models.Receta{}, id)
	return resultado.RowsAffected > 0
}
