package models

import "time"

type Maqueta struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UsuarioID   int    `json:"usuario_id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Escala      string `json:"escala"`
	Materiales  string `json:"materiales"`
	Dimensiones string `json:"dimensiones"`
}

type EvolucionMaqueta struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	MaquetaID   int       `json:"maqueta_id"`  // Relación con la maqueta
	Paso        int       `json:"paso"`        // Ej: 1, 2, 3...
	Titulo      string    `json:"titulo"`      // Ej: "Corte de base de MDF"
	Descripcion string    `json:"descripcion"` // Detalles del proceso
	Fecha       time.Time `json:"fecha"`
}
