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

	// Errores de autenticación
	ErrEmailEnUso            = errors.New("el correo ya esta registrado")
	ErrCredencialesInvalidas = errors.New("correo o contraseña incorrectos")
)
