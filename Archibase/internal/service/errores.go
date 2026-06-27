package service

import "errors"

// Definición de errores comunes de negocio
var (
	ErrDatosInvalidos          = errors.New("datos invalidos")
	ErrNombreObligatorio       = errors.New("el nombre es obligatorio")
	ErrEmailObligatorio        = errors.New("el correo electronico es obligatorio")
	ErrTituloObligatorio       = errors.New("el titulo es obligatorio")
	ErrIDMaquetaObligatorio    = errors.New("el ID de la maqueta es obligatorio")
	ErrTituloAvanceObligatorio = errors.New("el titulo del avance es obligatorio")
	ErrPasoInvalido            = errors.New("el numero de paso debe ser mayor a 0")

	// Errores de recursos no encontrados
	ErrUsuarioNoEncontrado   = errors.New("usuario no encontrado")
	ErrMaquetaNoEncontrada   = errors.New("maqueta no encontrada")
	ErrEvolucionNoEncontrada = errors.New("evolucion no encontrada")
	ErrRecetaNoEncontrada    = errors.New("receta no encontrada")

	// Errores de Proveedores
	ErrNombreProveedorObligatorio = errors.New("el nombre es obligatorio")
	ErrProveedorNoEncontrado      = errors.New("proveedor no encontrado")

	// Errores de Materiales
	ErrNombreMaterialObligatorio = errors.New("el nombre es obligatorio")
	ErrMaterialNoEncontrado      = errors.New("material no encontrado")

	// Errores de Ubicaciones
	ErrProvinciaUbicacionObligatoria = errors.New("la provincia es obligatoria")
	ErrCiudadUbicacionObligatoria    = errors.New("la ciudad es obligatoria")
	ErrUbicacionNoEncontrada         = errors.New("ubicacion no encontrada")
)
