package storage

import (
	"proyecto/internal/models"
	"testing"
)

// ==========================================
//     TESTS LEGACY (en memoria, sin BD)
//     Cubre las funciones globales de
//     asesor_storage.go
// ==========================================
//
// Nota: estas funciones usan variables globales de paquete
// (asesores, servicios, contrataciones), así que los tests no
// asumen listas vacías al inicio: comparan tamaños antes/después.

func TestLegacy_Asesor_CRUD(t *testing.T) {
	antes := len(GetAllAsesores())

	creado := CreateAsesor(models.Asesor{
		Nombre:       "Arq. Mendoza",
		Especialidad: "Diseño Interior",
		Experiencia:  "5 años",
		Contacto:     "0987654321",
		Modalidad:    "remoto",
	})
	if creado.IDasesor == 0 {
		t.Fatal("esperaba ID asignado al crear asesor (legacy)")
	}

	despues := len(GetAllAsesores())
	if despues != antes+1 {
		t.Fatalf("esperaba %d asesores tras crear, obtuve %d", antes+1, despues)
	}

	encontrado, err := GetAsesorByID(creado.IDasesor)
	if err != nil {
		t.Fatalf("esperaba encontrar asesor recién creado: %v", err)
	}
	if encontrado.Nombre != "Arq. Mendoza" {
		t.Fatalf("esperaba 'Arq. Mendoza', obtuve '%s'", encontrado.Nombre)
	}

	_, err = GetAsesorByID(-999)
	if err == nil {
		t.Fatal("no debía encontrar asesor con ID inexistente")
	}

	actualizado, err := UpdateAsesor(creado.IDasesor, models.Asesor{
		Nombre:       "Arq. Mendoza López",
		Especialidad: "Diseño Interior Avanzado",
		Experiencia:  "6 años",
		Contacto:     "0987654322",
		Modalidad:    "presencial",
	})
	if err != nil {
		t.Fatalf("esperaba actualizar asesor: %v", err)
	}
	if actualizado.Nombre != "Arq. Mendoza López" {
		t.Fatalf("esperaba 'Arq. Mendoza López', obtuve '%s'", actualizado.Nombre)
	}

	_, err = UpdateAsesor(-999, models.Asesor{Nombre: "X"})
	if err == nil {
		t.Fatal("no debía actualizar asesor inexistente")
	}

	if err := DeleteAsesor(creado.IDasesor); err != nil {
		t.Fatalf("esperaba eliminar asesor: %v", err)
	}
	if err := DeleteAsesor(-999); err == nil {
		t.Fatal("no debía eliminar asesor inexistente")
	}
}

func TestLegacy_Servicio_CRUD(t *testing.T) {
	antes := len(GetAllServicios())

	creado := CreateServicio(models.Servicio{
		Titulo:         "Consultoría rápida",
		Descripcion:    "Revisión express de anteproyecto",
		Precio:         50.00,
		Disponibilidad: "en línea",
		IDasesor:       1,
	})
	if creado.IDservicio == 0 {
		t.Fatal("esperaba ID asignado al crear servicio (legacy)")
	}

	despues := len(GetAllServicios())
	if despues != antes+1 {
		t.Fatalf("esperaba %d servicios tras crear, obtuve %d", antes+1, despues)
	}

	encontrado, err := GetServicioByID(creado.IDservicio)
	if err != nil {
		t.Fatalf("esperaba encontrar servicio recién creado: %v", err)
	}
	if encontrado.Titulo != "Consultoría rápida" {
		t.Fatalf("esperaba 'Consultoría rápida', obtuve '%s'", encontrado.Titulo)
	}

	_, err = GetServicioByID(-999)
	if err == nil {
		t.Fatal("no debía encontrar servicio con ID inexistente")
	}

	actualizado, err := UpdateServicio(creado.IDservicio, models.Servicio{
		Titulo:         "Consultoría rápida Plus",
		Descripcion:    "Revisión express + informe",
		Precio:         70.00,
		Disponibilidad: "en línea y presencial",
		IDasesor:       2,
	})
	if err != nil {
		t.Fatalf("esperaba actualizar servicio: %v", err)
	}
	if actualizado.Titulo != "Consultoría rápida Plus" {
		t.Fatalf("esperaba 'Consultoría rápida Plus', obtuve '%s'", actualizado.Titulo)
	}

	_, err = UpdateServicio(-999, models.Servicio{Titulo: "X"})
	if err == nil {
		t.Fatal("no debía actualizar servicio inexistente")
	}

	if err := DeleteServicio(creado.IDservicio); err != nil {
		t.Fatalf("esperaba eliminar servicio: %v", err)
	}
	if err := DeleteServicio(-999); err == nil {
		t.Fatal("no debía eliminar servicio inexistente")
	}
}

func TestLegacy_Contratacion_CRUD(t *testing.T) {
	antes := len(GetAllContrataciones())

	creado := CreateContratacion(models.Contratacion{
		Estudiante: "Luis Zambrano",
		Fecha:      "2026-07-08",
		Estado:     "pendiente",
		IDservicio: 1,
	})
	if creado.IDcontratacion == 0 {
		t.Fatal("esperaba ID asignado al crear contratación (legacy)")
	}

	despues := len(GetAllContrataciones())
	if despues != antes+1 {
		t.Fatalf("esperaba %d contrataciones tras crear, obtuve %d", antes+1, despues)
	}

	encontrada, err := GetContratacionByID(creado.IDcontratacion)
	if err != nil {
		t.Fatalf("esperaba encontrar contratación recién creada: %v", err)
	}
	if encontrada.Estudiante != "Luis Zambrano" {
		t.Fatalf("esperaba 'Luis Zambrano', obtuve '%s'", encontrada.Estudiante)
	}

	_, err = GetContratacionByID(-999)
	if err == nil {
		t.Fatal("no debía encontrar contratación con ID inexistente")
	}

	actualizada, err := UpdateContratacion(creado.IDcontratacion, models.Contratacion{
		Estudiante: "Luis Zambrano",
		Fecha:      "2026-07-10",
		Estado:     "confirmada",
		IDservicio: 1,
	})
	if err != nil {
		t.Fatalf("esperaba actualizar contratación: %v", err)
	}
	if actualizada.Estado != "confirmada" {
		t.Fatalf("esperaba estado 'confirmada', obtuve '%s'", actualizada.Estado)
	}

	_, err = UpdateContratacion(-999, models.Contratacion{Estudiante: "X"})
	if err == nil {
		t.Fatal("no debía actualizar contratación inexistente")
	}

	if err := DeleteContratacion(creado.IDcontratacion); err != nil {
		t.Fatalf("esperaba eliminar contratación: %v", err)
	}
	if err := DeleteContratacion(-999); err == nil {
		t.Fatal("no debía eliminar contratación inexistente")
	}
}
