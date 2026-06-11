package storage

import (
	"errors"
	"proyecto/internal/models"
)

// ---- ASESORES ----
var asesores []models.Asesor
var nextAsesorID = 1

func GetAllAsesores() []models.Asesor {
	return asesores
}

func GetAsesorByID(id int) (models.Asesor, error) {
	for _, a := range asesores {
		if a.IDasesor == id {
			return a, nil
		}
	}
	return models.Asesor{}, errors.New("asesor no encontrado")
}

func CreateAsesor(a models.Asesor) models.Asesor {
	a.IDasesor = nextAsesorID
	nextAsesorID++
	asesores = append(asesores, a)
	return a
}

func UpdateAsesor(id int, nuevo models.Asesor) (models.Asesor, error) {
	for i, a := range asesores {
		if a.IDasesor == id {
			nuevo.IDasesor = id
			asesores[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Asesor{}, errors.New("asesor no encontrado")
}

func DeleteAsesor(id int) error {
	for i, a := range asesores {
		if a.IDasesor == id {
			asesores = append(asesores[:i], asesores[i+1:]...)
			return nil
		}
	}
	return errors.New("asesor no encontrado")
}
