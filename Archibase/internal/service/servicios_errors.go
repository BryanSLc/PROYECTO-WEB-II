package service

import "errors"

// Errores específicos para Servicio y Asesor.
// (Se separan para evitar conflictos si el archivo errores.go cambia.)
var (
	ErrServicioNoEncontrado = errors.New("servicio no encontrado")
	ErrAsesorNoEncontrado   = errors.New("asesor no encontrado")
)
