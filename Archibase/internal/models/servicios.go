package models

type Asesor struct {
	IDasesor     int    `json:"id_asesor"`
	Nombre       string `json:"nombre"`
	Especialidad string `json:"especialidad"`
	Experiencia  string `json:"experiencia"`
	Contacto     string `json:"contacto"`
	Modalidad    string `json:"modalidad"` // "presencial", "virtual", "ambas"
}

type Servicio struct {
	IDservicio     int     `json:"id_servicio"`
	Titulo         string  `json:"titulo"`
	Descripcion    string  `json:"descripcion"`
	Precio         float64 `json:"precio"`
	Disponibilidad string  `json:"disponibilidad"`
	IDasesor       int     `json:"id_asesor"`
}

type Contratacion struct {
	IDcontratacion int    `json:"id_contratacion"`
	Estudiante     string `json:"estudiante"`
	Fecha          string `json:"fecha"`
	Estado         string `json:"estado"` // "pendiente", "confirmada", "cancelada","completado"
	IDservicio     int    `json:"id_servicio"`
}
