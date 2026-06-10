# PROYECTO-WEB-2

estructuras

========================
DIAGRAMA ER DEL SISTEMA
========================


┌──────────────────────┐
│       USUARIO        │
├──────────────────────┤
│ id_usuario (PK)      │
│ nombre               │
│ correo               │
│ contraseña           │
│ universidad          │
│ ciudad               │
│ fecha_registro       │
└──────────┬───────────┘
           │
           │ 1 usuario publica muchas
           │
           ▼
┌──────────────────────┐
│       MAQUETA        │
├──────────────────────┤
│ id_maqueta (PK)      │
│ titulo               │
│ descripcion          │
│ escala               │
│ categoria            │
│ dificultad           │
│ tiempo_estimado      │
│ imagen               │
│ id_usuario (FK)      │
└──────────┬───────────┘
           │
           │ 1 maqueta contiene muchos
           │
           ▼
┌──────────────────────┐
│      MATERIAL        │
├──────────────────────┤
│ id_material (PK)     │
│ nombre               │
│ descripcion          │
│ tipo                 │
│ costo_aproximado     │
│ cantidad_aproximada  │
│ id_maqueta (FK)      │
└──────────┬───────────┘
           │
           │ muchos materiales pueden
           │ encontrarse en muchos
           │ proveedores
           │
           ▼
┌──────────────────────┐
│     PROVEEDOR        │
├──────────────────────┤
│ id_proveedor (PK)    │
│ nombre               │
│ ciudad               │
│ provincia            │
│ direccion            │
│ telefono             │
│ horario              │
│ tipo_material        │
└──────────────────────┘