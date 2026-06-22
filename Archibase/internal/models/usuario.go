package models

import "time"

type Usuario struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	Semestre      int       `json:"semestre"`
	Telefono      string    `json:"telefono"`
	Rol           string    `json:"rol"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}
