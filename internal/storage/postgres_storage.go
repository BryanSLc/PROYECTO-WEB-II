package storage

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"proyecto/internal/models"
)

// PostgresStorage es el storage de producción (Docker).
// Los tests siguen usando SQLiteStorage y Memoria — este archivo no los afecta.
type PostgresStorage struct {
	db *gorm.DB
}

func NuevoPostgresStorage(dsn string) *PostgresStorage {
	var db *gorm.DB
	var err error

	maxIntentos := 10
	for intento := 1; intento <= maxIntentos; intento++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("No se pudo conectar a Postgres (intento %d/%d): %v", intento, maxIntentos, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Error al conectar a Postgres tras %d intentos: %v", maxIntentos, err)
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Maqueta{},
		&models.EvolucionMaqueta{},
		&models.Receta{},
		&models.Proveedor{},
		&models.MaterialProveedor{},
		&models.Ubicacion{},
		&models.Asesor{},
		&models.Servicio{},
		&models.Contratacion{},
	)
	if err != nil {
		log.Fatalf("Error al realizar la migración: %v", err)
	}

	return &PostgresStorage{db: db}
}

func ConstruirDSN(host, port, user, password, dbname, sslmode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
}

// ==========================================
//          MÉTODOS PARA PROVEEDORES
// ==========================================

func (s *PostgresStorage) CrearProveedor(proveedor models.Proveedor) models.Proveedor {
	s.db.Create(&proveedor)
	return proveedor
}

func (s *PostgresStorage) ListarProveedores() []models.Proveedor {
	var proveedores []models.Proveedor
	s.db.Find(&proveedores)
	return proveedores
}

func (s *PostgresStorage) BuscarProveedorPorID(id int) (models.Proveedor, bool) {
	var proveedor models.Proveedor
	if err := s.db.First(&proveedor, id).Error; err != nil {
		return models.Proveedor{}, false
	}
	return proveedor, true
}

func (s *PostgresStorage) ActualizarProveedor(id int, datos models.Proveedor) (models.Proveedor, bool) {
	var proveedor models.Proveedor
	if err := s.db.First(&proveedor, id).Error; err != nil {
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

func (s *PostgresStorage) EliminarProveedor(id int) bool {
	return s.db.Delete(&models.Proveedor{}, id).RowsAffected > 0
}

// ==========================================
//     MÉTODOS PARA MATERIAL PROVEEDOR
// ==========================================

func (s *PostgresStorage) CrearMaterial(material models.MaterialProveedor) models.MaterialProveedor {
	s.db.Create(&material)
	return material
}

func (s *PostgresStorage) ListarMateriales() []models.MaterialProveedor {
	var materiales []models.MaterialProveedor
	s.db.Find(&materiales)
	return materiales
}

func (s *PostgresStorage) BuscarMaterialPorID(id int) (models.MaterialProveedor, bool) {
	var material models.MaterialProveedor
	if err := s.db.First(&material, id).Error; err != nil {
		return models.MaterialProveedor{}, false
	}
	return material, true
}

func (s *PostgresStorage) ActualizarMaterial(id int, datos models.MaterialProveedor) (models.MaterialProveedor, bool) {
	var material models.MaterialProveedor
	if err := s.db.First(&material, id).Error; err != nil {
		return models.MaterialProveedor{}, false
	}
	material.Nombre = datos.Nombre
	material.Categoria = datos.Categoria
	material.PrecioReferencial = datos.PrecioReferencial
	material.IDProveedor = datos.IDProveedor
	s.db.Save(&material)
	return material, true
}

func (s *PostgresStorage) EliminarMaterial(id int) bool {
	return s.db.Delete(&models.MaterialProveedor{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA UBICACIONES
// ==========================================

func (s *PostgresStorage) CrearUbicacion(ubicacion models.Ubicacion) models.Ubicacion {
	s.db.Create(&ubicacion)
	return ubicacion
}

func (s *PostgresStorage) ListarUbicaciones() []models.Ubicacion {
	var ubicaciones []models.Ubicacion
	s.db.Find(&ubicaciones)
	return ubicaciones
}

func (s *PostgresStorage) BuscarUbicacionPorID(id int) (models.Ubicacion, bool) {
	var ubicacion models.Ubicacion
	if err := s.db.First(&ubicacion, id).Error; err != nil {
		return models.Ubicacion{}, false
	}
	return ubicacion, true
}

func (s *PostgresStorage) ActualizarUbicacion(id int, datos models.Ubicacion) (models.Ubicacion, bool) {
	var ubicacion models.Ubicacion
	if err := s.db.First(&ubicacion, id).Error; err != nil {
		return models.Ubicacion{}, false
	}
	ubicacion.Provincia = datos.Provincia
	ubicacion.Ciudad = datos.Ciudad
	s.db.Save(&ubicacion)
	return ubicacion, true
}

func (s *PostgresStorage) EliminarUbicacion(id int) bool {
	return s.db.Delete(&models.Ubicacion{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA MAQUETAS
// ==========================================

func (s *PostgresStorage) CrearMaqueta(maqueta models.Maqueta) models.Maqueta {
	s.db.Create(&maqueta)
	return maqueta
}

func (s *PostgresStorage) ListarMaquetas() []models.Maqueta {
	var maquetas []models.Maqueta
	s.db.Find(&maquetas)
	return maquetas
}

func (s *PostgresStorage) BuscarMaquetaPorID(id int) (models.Maqueta, bool) {
	var maqueta models.Maqueta
	if err := s.db.First(&maqueta, id).Error; err != nil {
		return models.Maqueta{}, false
	}
	return maqueta, true
}

func (s *PostgresStorage) ActualizarMaqueta(id int, datos models.Maqueta) (models.Maqueta, bool) {
	var maqueta models.Maqueta
	if err := s.db.First(&maqueta, id).Error; err != nil {
		return models.Maqueta{}, false
	}
	maqueta.UsuarioID = datos.UsuarioID
	maqueta.Titulo = datos.Titulo
	maqueta.Descripcion = datos.Descripcion
	maqueta.Escala = datos.Escala
	maqueta.Materiales = datos.Materiales
	maqueta.Dimensiones = datos.Dimensiones
	s.db.Save(&maqueta)
	return maqueta, true
}

func (s *PostgresStorage) EliminarMaqueta(id int) bool {
	return s.db.Delete(&models.Maqueta{}, id).RowsAffected > 0
}

// ==========================================
//       MÉTODOS PARA EVOLUCIÓN MAQUETA
// ==========================================

func (s *PostgresStorage) AgregarEvolucion(evolucion models.EvolucionMaqueta) models.EvolucionMaqueta {
	evolucion.Fecha = time.Now()
	s.db.Create(&evolucion)
	return evolucion
}

func (s *PostgresStorage) ListarEvolucionPorMaqueta(maquetaID int) []models.EvolucionMaqueta {
	var historial []models.EvolucionMaqueta
	s.db.Where("maqueta_id = ?", maquetaID).Find(&historial)
	return historial
}

func (s *PostgresStorage) EliminarEvolucion(id int) bool {
	return s.db.Delete(&models.EvolucionMaqueta{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA USUARIOS
// ==========================================

func (s *PostgresStorage) CrearUsuario(usuario models.Usuario) models.Usuario {
	usuario.FechaCreacion = time.Now()
	s.db.Create(&usuario)
	return usuario
}

func (s *PostgresStorage) ListarUsuarios() []models.Usuario {
	var usuarios []models.Usuario
	s.db.Find(&usuarios)
	return usuarios
}

func (s *PostgresStorage) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var usuario models.Usuario
	if err := s.db.First(&usuario, id).Error; err != nil {
		return models.Usuario{}, false
	}
	return usuario, true
}

func (s *PostgresStorage) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var usuario models.Usuario
	if err := s.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return models.Usuario{}, false
	}
	return usuario, true
}

func (s *PostgresStorage) ActualizarUsuario(id int, datos models.Usuario) (models.Usuario, bool) {
	var usuario models.Usuario
	if err := s.db.First(&usuario, id).Error; err != nil {
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

func (s *PostgresStorage) EliminarUsuario(id int) bool {
	return s.db.Delete(&models.Usuario{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA RECETAS
// ==========================================

func (s *PostgresStorage) CrearReceta(receta models.Receta) models.Receta {
	s.db.Create(&receta)
	return receta
}

func (s *PostgresStorage) ListarRecetas() []models.Receta {
	var recetas []models.Receta
	s.db.Find(&recetas)
	return recetas
}

func (s *PostgresStorage) ListarRecetasPorMaqueta(maquetaID int) []models.Receta {
	var recetas []models.Receta
	s.db.Where("maqueta_id = ?", maquetaID).Find(&recetas)
	return recetas
}

func (s *PostgresStorage) BuscarRecetaPorID(id int) (models.Receta, bool) {
	var receta models.Receta
	if err := s.db.First(&receta, id).Error; err != nil {
		return models.Receta{}, false
	}
	return receta, true
}

func (s *PostgresStorage) ActualizarReceta(id int, datos models.Receta) (models.Receta, bool) {
	var receta models.Receta
	if err := s.db.First(&receta, id).Error; err != nil {
		return models.Receta{}, false
	}
	receta.MaquetaID = datos.MaquetaID
	receta.Titulo = datos.Titulo
	receta.Descripcion = datos.Descripcion
	receta.Pasos = datos.Pasos
	s.db.Save(&receta)
	return receta, true
}

func (s *PostgresStorage) EliminarReceta(id int) bool {
	return s.db.Delete(&models.Receta{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA ASESORES
// ==========================================

func (s *PostgresStorage) CrearAsesor(a models.Asesor) models.Asesor {
	s.db.Create(&a)
	return a
}

func (s *PostgresStorage) ListarAsesores() []models.Asesor {
	var lista []models.Asesor
	s.db.Find(&lista)
	return lista
}

func (s *PostgresStorage) BuscarAsesorPorID(id int) (models.Asesor, bool) {
	var a models.Asesor
	if err := s.db.First(&a, id).Error; err != nil {
		return models.Asesor{}, false
	}
	return a, true
}

func (s *PostgresStorage) ActualizarAsesor(id int, datos models.Asesor) (models.Asesor, bool) {
	var a models.Asesor
	if err := s.db.First(&a, id).Error; err != nil {
		return models.Asesor{}, false
	}
	a.Nombre = datos.Nombre
	a.Especialidad = datos.Especialidad
	a.Experiencia = datos.Experiencia
	a.Contacto = datos.Contacto
	a.Modalidad = datos.Modalidad
	s.db.Save(&a)
	return a, true
}

func (s *PostgresStorage) EliminarAsesor(id int) bool {
	return s.db.Delete(&models.Asesor{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA CONTRATACION
// ==========================================

func (s *PostgresStorage) CrearContratacion(c models.Contratacion) models.Contratacion {
	s.db.Create(&c)
	return c
}

func (s *PostgresStorage) ListarContrataciones() []models.Contratacion {
	var lista []models.Contratacion
	s.db.Find(&lista)
	return lista
}

func (s *PostgresStorage) BuscarContratacionPorID(id int) (models.Contratacion, bool) {
	var c models.Contratacion
	if err := s.db.First(&c, id).Error; err != nil {
		return models.Contratacion{}, false
	}
	return c, true
}

func (s *PostgresStorage) ActualizarContratacion(id int, datos models.Contratacion) (models.Contratacion, bool) {
	var c models.Contratacion
	if err := s.db.First(&c, id).Error; err != nil {
		return models.Contratacion{}, false
	}
	c.Estudiante = datos.Estudiante
	c.Fecha = datos.Fecha
	c.Estado = datos.Estado
	c.IDservicio = datos.IDservicio
	s.db.Save(&c)
	return c, true
}

func (s *PostgresStorage) EliminarContratacion(id int) bool {
	return s.db.Delete(&models.Contratacion{}, id).RowsAffected > 0
}

// ==========================================
//          MÉTODOS PARA SERVICIO
// ==========================================

func (s *PostgresStorage) CrearServicio(serv models.Servicio) models.Servicio {
	s.db.Create(&serv)
	return serv
}

func (s *PostgresStorage) ListarServicios() []models.Servicio {
	var lista []models.Servicio
	s.db.Find(&lista)
	return lista
}

func (s *PostgresStorage) BuscarServicioPorID(id int) (models.Servicio, bool) {
	var serv models.Servicio
	if err := s.db.First(&serv, id).Error; err != nil {
		return models.Servicio{}, false
	}
	return serv, true
}

func (s *PostgresStorage) ActualizarServicio(id int, datos models.Servicio) (models.Servicio, bool) {
	var serv models.Servicio
	if err := s.db.First(&serv, id).Error; err != nil {
		return models.Servicio{}, false
	}
	serv.Titulo = datos.Titulo
	serv.Descripcion = datos.Descripcion
	serv.Precio = datos.Precio
	serv.Disponibilidad = datos.Disponibilidad
	serv.IDasesor = datos.IDasesor
	s.db.Save(&serv)
	return serv, true
}

func (s *PostgresStorage) EliminarServicio(id int) bool {
	return s.db.Delete(&models.Servicio{}, id).RowsAffected > 0
}
