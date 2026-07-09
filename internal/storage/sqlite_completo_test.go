package storage

import (
	"proyecto/internal/models"
	"testing"
)

func nuevoAlmacenTest(t *testing.T) *SQLiteStorage {
	t.Helper()
	return NuevoSQLiteStorage("file::memory:?cache=shared&_pragma=foreign_keys(1)")
}

func TestSQLite_Maqueta_CRUD(t *testing.T) {
	a := nuevoAlmacenTest(t)

	mq := a.CrearMaqueta(models.Maqueta{Titulo: "Casa A", UsuarioID: 1})
	if mq.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	lista := a.ListarMaquetas()
	if len(lista) == 0 {
		t.Fatal("esperaba al menos 1 maqueta")
	}

	encontrada, ok := a.BuscarMaquetaPorID(mq.ID)
	if !ok {
		t.Fatal("esperaba encontrar maqueta")
	}
	if encontrada.Titulo != "Casa A" {
		t.Fatalf("esperaba 'Casa A', obtuve '%s'", encontrada.Titulo)
	}

	_, ok = a.BuscarMaquetaPorID(999)
	if ok {
		t.Fatal("no debía encontrar maqueta inexistente")
	}

	actualizada, ok := a.ActualizarMaqueta(mq.ID, models.Maqueta{Titulo: "Casa B"})
	if !ok {
		t.Fatal("esperaba actualizar maqueta")
	}
	if actualizada.Titulo != "Casa B" {
		t.Fatalf("esperaba 'Casa B', obtuve '%s'", actualizada.Titulo)
	}

	_, ok = a.ActualizarMaqueta(999, models.Maqueta{Titulo: "X"})
	if ok {
		t.Fatal("no debía actualizar maqueta inexistente")
	}

	if !a.EliminarMaqueta(mq.ID) {
		t.Fatal("esperaba eliminar maqueta")
	}
	if a.EliminarMaqueta(999) {
		t.Fatal("no debía eliminar maqueta inexistente")
	}
}

func TestSQLite_Evolucion_CRUD(t *testing.T) {
	a := nuevoAlmacenTest(t)

	mq := a.CrearMaqueta(models.Maqueta{Titulo: "Casa A"})

	e := a.AgregarEvolucion(models.EvolucionMaqueta{MaquetaID: mq.ID, Titulo: "Avance 1", Paso: 1})
	if e.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	historial := a.ListarEvolucionPorMaqueta(mq.ID)
	if len(historial) == 0 {
		t.Fatal("esperaba al menos 1 evolución")
	}

	if !a.EliminarEvolucion(e.ID) {
		t.Fatal("esperaba eliminar evolución")
	}
	if a.EliminarEvolucion(999) {
		t.Fatal("no debía eliminar evolución inexistente")
	}
}

func TestSQLite_Receta_CRUD(t *testing.T) {
	a := nuevoAlmacenTest(t)

	r := a.CrearReceta(models.Receta{Titulo: "Receta A", MaquetaID: 1})
	if r.ID == 0 {
		t.Fatal("esperaba ID asignado")
	}

	lista := a.ListarRecetas()
	if len(lista) == 0 {
		t.Fatal("esperaba al menos 1 receta")
	}

	porMaqueta := a.ListarRecetasPorMaqueta(1)
	if len(porMaqueta) == 0 {
		t.Fatal("esperaba al menos 1 receta por maqueta")
	}

	encontrada, ok := a.BuscarRecetaPorID(r.ID)
	if !ok {
		t.Fatal("esperaba encontrar receta")
	}
	if encontrada.Titulo != "Receta A" {
		t.Fatalf("esperaba 'Receta A', obtuve '%s'", encontrada.Titulo)
	}

	_, ok = a.BuscarRecetaPorID(999)
	if ok {
		t.Fatal("no debía encontrar receta inexistente")
	}

	actualizada, ok := a.ActualizarReceta(r.ID, models.Receta{Titulo: "Receta B", MaquetaID: 1})
	if !ok {
		t.Fatal("esperaba actualizar receta")
	}
	if actualizada.Titulo != "Receta B" {
		t.Fatalf("esperaba 'Receta B', obtuve '%s'", actualizada.Titulo)
	}

	_, ok = a.ActualizarReceta(999, models.Receta{Titulo: "X"})
	if ok {
		t.Fatal("no debía actualizar receta inexistente")
	}

	if !a.EliminarReceta(r.ID) {
		t.Fatal("esperaba eliminar receta")
	}
	if a.EliminarReceta(999) {
		t.Fatal("no debía eliminar receta inexistente")
	}
}

func TestSQLite_Proveedor_ActualizarEliminar(t *testing.T) {
	a := nuevoAlmacenTest(t)

	p := a.CrearProveedor(models.Proveedor{Nombre: "Ferretería A", Ciudad: "Manta"})

	actualizado, ok := a.ActualizarProveedor(p.ID, models.Proveedor{Nombre: "Ferretería B", Ciudad: "Portoviejo"})
	if !ok {
		t.Fatal("esperaba actualizar proveedor")
	}
	if actualizado.Nombre != "Ferretería B" {
		t.Fatalf("esperaba 'Ferretería B', obtuve '%s'", actualizado.Nombre)
	}

	_, ok = a.ActualizarProveedor(999, models.Proveedor{Nombre: "X"})
	if ok {
		t.Fatal("no debía actualizar proveedor inexistente")
	}

	if !a.EliminarProveedor(p.ID) {
		t.Fatal("esperaba eliminar proveedor")
	}
	if a.EliminarProveedor(999) {
		t.Fatal("no debía eliminar proveedor inexistente")
	}
}

func TestSQLite_Material_ActualizarEliminar(t *testing.T) {
	a := nuevoAlmacenTest(t)

	mat := a.CrearMaterial(models.MaterialProveedor{Nombre: "MDF", Categoria: "Madera"})

	actualizado, ok := a.ActualizarMaterial(mat.ID, models.MaterialProveedor{Nombre: "MDF 15mm", Categoria: "Madera"})
	if !ok {
		t.Fatal("esperaba actualizar material")
	}
	if actualizado.Nombre != "MDF 15mm" {
		t.Fatalf("esperaba 'MDF 15mm', obtuve '%s'", actualizado.Nombre)
	}

	_, ok = a.ActualizarMaterial(999, models.MaterialProveedor{Nombre: "X"})
	if ok {
		t.Fatal("no debía actualizar material inexistente")
	}

	if !a.EliminarMaterial(mat.ID) {
		t.Fatal("esperaba eliminar material")
	}
	if a.EliminarMaterial(999) {
		t.Fatal("no debía eliminar material inexistente")
	}
}

func TestSQLite_Ubicacion_ActualizarEliminar(t *testing.T) {
	a := nuevoAlmacenTest(t)

	u := a.CrearUbicacion(models.Ubicacion{Provincia: "Manabí", Ciudad: "Manta"})

	actualizada, ok := a.ActualizarUbicacion(u.ID, models.Ubicacion{Provincia: "Guayas", Ciudad: "Guayaquil"})
	if !ok {
		t.Fatal("esperaba actualizar ubicacion")
	}
	if actualizada.Ciudad != "Guayaquil" {
		t.Fatalf("esperaba 'Guayaquil', obtuve '%s'", actualizada.Ciudad)
	}

	_, ok = a.ActualizarUbicacion(999, models.Ubicacion{Provincia: "X"})
	if ok {
		t.Fatal("no debía actualizar ubicacion inexistente")
	}

	if !a.EliminarUbicacion(u.ID) {
		t.Fatal("esperaba eliminar ubicacion")
	}
	if a.EliminarUbicacion(999) {
		t.Fatal("no debía eliminar ubicacion inexistente")
	}
}

func TestSQLite_Asesor_ActualizarEliminar(t *testing.T) {
	a := nuevoAlmacenTest(t)

	as := a.CrearAsesor(models.Asesor{Nombre: "Arq. Pérez", Especialidad: "Diseño"})

	actualizado, ok := a.ActualizarAsesor(as.IDasesor, models.Asesor{Nombre: "Arq. López", Especialidad: "Estructuras"})
	if !ok {
		t.Fatal("esperaba actualizar asesor")
	}
	if actualizado.Nombre != "Arq. López" {
		t.Fatalf("esperaba 'Arq. López', obtuve '%s'", actualizado.Nombre)
	}

	_, ok = a.ActualizarAsesor(999, models.Asesor{Nombre: "X"})
	if ok {
		t.Fatal("no debía actualizar asesor inexistente")
	}

	if !a.EliminarAsesor(as.IDasesor) {
		t.Fatal("esperaba eliminar asesor")
	}
	if a.EliminarAsesor(999) {
		t.Fatal("no debía eliminar asesor inexistente")
	}
}

func TestSQLite_Usuario_ActualizarEliminar(t *testing.T) {
	a := nuevoAlmacenTest(t)

	u := a.CrearUsuario(models.Usuario{Nombre: "Ana", Email: "ana@test.com"})

	encontrado, ok := a.BuscarUsuarioPorID(u.ID)
	if !ok {
		t.Fatal("esperaba encontrar usuario por ID")
	}
	if encontrado.Nombre != "Ana" {
		t.Fatalf("esperaba 'Ana', obtuve '%s'", encontrado.Nombre)
	}

	_, ok = a.BuscarUsuarioPorID(999)
	if ok {
		t.Fatal("no debía encontrar usuario inexistente")
	}

	actualizado, ok := a.ActualizarUsuario(u.ID, models.Usuario{Nombre: "Ana Lopez", Email: "ana@test.com"})
	if !ok {
		t.Fatal("esperaba actualizar usuario")
	}
	if actualizado.Nombre != "Ana Lopez" {
		t.Fatalf("esperaba 'Ana Lopez', obtuve '%s'", actualizado.Nombre)
	}

	_, ok = a.ActualizarUsuario(999, models.Usuario{Nombre: "X"})
	if ok {
		t.Fatal("no debía actualizar usuario inexistente")
	}

	if !a.EliminarUsuario(u.ID) {
		t.Fatal("esperaba eliminar usuario")
	}
	if a.EliminarUsuario(999) {
		t.Fatal("no debía eliminar usuario inexistente")
	}
}
