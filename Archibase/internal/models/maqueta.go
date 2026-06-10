package models

type Maqueta struct {
	ID          int    `json:"id"`
	UsuarioID   int    `json:"usuario_id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Escala      string `json:"escala"`
	Materiales  string `json:"materiales"`
	Dimensiones string `json:"dimensiones"`
}
