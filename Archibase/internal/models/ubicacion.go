package models

type Ubicacion struct {
	ID        int    `json:"id"`
	Provincia string `json:"provincia"`
	Ciudad    string `json:"ciudad"`
}
