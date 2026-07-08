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

// ---- SERVICIOS ----
var servicios []models.Servicio
var nextServicioID = 1

func GetAllServicios() []models.Servicio {
	return servicios
}

func GetServicioByID(id int) (models.Servicio, error) {
	for _, s := range servicios {
		if s.IDservicio == id {
			return s, nil
		}
	}
	return models.Servicio{}, errors.New("servicio no encontrado")
}

func CreateServicio(s models.Servicio) models.Servicio {
	s.IDservicio = nextServicioID
	nextServicioID++
	servicios = append(servicios, s)
	return s
}

func UpdateServicio(id int, nuevo models.Servicio) (models.Servicio, error) {
	for i, s := range servicios {
		if s.IDservicio == id {
			nuevo.IDservicio = id
			servicios[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Servicio{}, errors.New("servicio no encontrado")
}

func DeleteServicio(id int) error {
	for i, s := range servicios {
		if s.IDservicio == id {
			servicios = append(servicios[:i], servicios[i+1:]...)
			return nil
		}
	}
	return errors.New("servicio no encontrado")
}

// ---- CONTRATACIONES ----
var contrataciones []models.Contratacion
var nextContratacionID = 1

func GetAllContrataciones() []models.Contratacion {
	return contrataciones
}

func GetContratacionByID(id int) (models.Contratacion, error) {
	for _, c := range contrataciones {
		if c.IDcontratacion == id {
			return c, nil
		}
	}
	return models.Contratacion{}, errors.New("contratación no encontrada")
}

func CreateContratacion(c models.Contratacion) models.Contratacion {
	c.IDcontratacion = nextContratacionID
	nextContratacionID++
	contrataciones = append(contrataciones, c)
	return c
}

func UpdateContratacion(id int, nuevo models.Contratacion) (models.Contratacion, error) {
	for i, c := range contrataciones {
		if c.IDcontratacion == id {
			nuevo.IDcontratacion = id
			contrataciones[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Contratacion{}, errors.New("contratación no encontrada")
}

func DeleteContratacion(id int) error {
	for i, c := range contrataciones {
		if c.IDcontratacion == id {
			contrataciones = append(contrataciones[:i], contrataciones[i+1:]...)
			return nil
		}
	}
	return errors.New("contratación no encontrada")
}
