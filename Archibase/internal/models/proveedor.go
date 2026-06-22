package models

type Proveedor struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Ciudad    string `json:"ciudad"`
	Provincia string `json:"provincia"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
