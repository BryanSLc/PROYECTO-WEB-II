package models

type Contratacion struct {
	IDcontratacion int    `json:"id_contratacion"`
	Estudiante     string `json:"estudiante"`
	Fecha          string `json:"fecha"`
	Estado         string `json:"estado"` // "pendiente", "confirmada", "cancelada","completado"
	IDservicio     int    `json:"id_servicio"`
}
