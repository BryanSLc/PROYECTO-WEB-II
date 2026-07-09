package models

type Servicio struct {
	IDservicio     int     `json:"id_servicio" gorm:"primaryKey;autoIncrement"`
	Titulo         string  `json:"titulo"`
	Descripcion    string  `json:"descripcion"`
	Precio         float64 `json:"precio"`
	Disponibilidad string  `json:"disponibilidad"`
	IDasesor       int     `json:"id_asesor"`
}
