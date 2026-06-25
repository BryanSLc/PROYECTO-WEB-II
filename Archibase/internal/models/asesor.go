package models

type Asesor struct {
	IDasesor     int    `json:"id_asesor"`
	Nombre       string `json:"nombre"`
	Especialidad string `json:"especialidad"`
	Experiencia  string `json:"experiencia"`
	Contacto     string `json:"contacto"`
	Modalidad    string `json:"modalidad"`
}