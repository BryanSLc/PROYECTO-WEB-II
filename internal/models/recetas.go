package models

type Paso struct {
	NumeroPaso  int    `json:"numero_paso"`
	Titulo      string `json:"titulo"`
	Instruccion string `json:"instruccion"`
	Calculo     string `json:"calculo"`
}

type Receta struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	MaquetaID   int    `json:"maqueta_id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Pasos       []Paso `json:"pasos" gorm:"type:text;serializer:json"`
}
