package models

type MaterialProveedor struct {
	ID                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Categoria         string  `json:"categoria"`
	PrecioReferencial float64 `json:"precio_referencial"`
	IDProveedor       int     `json:"id_proveedor"`
}
