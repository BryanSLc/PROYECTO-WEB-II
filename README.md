# Archibase

**Archibase** es una plataforma pensada para estudiantes de arquitectura que necesitan documentar, construir y conseguir recursos para sus maquetas (modelos a escala). La idea central es conectar todo el ciclo de una maqueta en un solo lugar:

- Un estudiante (**Usuario**) publica sus **Maquetas**, con su escala, categoría y dificultad.
- Cada maqueta tiene una **Receta**: los pasos de construcción, con instrucciones y cálculos, y un historial de **Evolución** (fotos/avances por etapa).
- Las maquetas necesitan **Materiales**, y cada material se puede conseguir en distintos **Proveedores**, ubicados en distintas **Ubicaciones** (ciudad/provincia).
- Si un estudiante necesita ayuda profesional, puede contactar a un **Asesor**, contratar uno de sus **Servicios**, y queda registrado como una **Contratación**.

En resumen: Archibase resuelve tres preguntas para un estudiante de arquitectura — *¿cómo lo construyo?*, *¿dónde consigo los materiales?* y *¿quién me puede asesorar?*.

## Equipo y responsabilidades

| Integrante | Módulo(s) a cargo |
|---|---|
| **Bryan Lopez** | Usuario / Autenticación (JWT), Maqueta + Evolución, Receta |
| **Eduardo García** | Asesor, Servicio, Contratación |
| **isaac** | Proveedor, Ubicación, Material |

## Stack técnico

| Capa | Tecnología |
|---|---|
| Lenguaje | Go 1.25 |
| Router HTTP | [chi](https://github.com/go-chi/chi) v5 |
| ORM / Persistencia | [GORM](https://gorm.io) sobre PostgreSQL 16 |
| Autenticación | JWT ([golang-jwt/jwt](https://github.com/golang-jwt/jwt)) + hashing con `bcrypt` |
| Base de datos (prod) | PostgreSQL 16 (Alpine) |
| Contenedores | Docker multi-stage (`golang:1.25-alpine` → `alpine:3.20`), Docker Compose |
| CI/CD | GitHub Actions (`go vet` → `go build` → `go test` con cobertura) |
| Tests | `testing` estándar de Go, con mocks propios de los repositorios |

## Arquitectura

El proyecto sigue una arquitectura en capas:

```
HTTP Request
     │
     ▼
 Handler (internal/handlers)   → decodifica el request, valida formato, arma la respuesta
     │
     ▼
 Service (internal/service)    → reglas de negocio, validaciones de dominio
     │
     ▼
 Storage / Repository (internal/storage) → acceso a datos vía GORM (PostgresStorage)
     │
     ▼
 PostgreSQL
```

Cada servicio (`UsuarioService`, `MaquetaService`, `RecetaService`, `AuthService`, etc.) define su propia interfaz de repositorio y la recibe por inyección de dependencias desde `cmd/api/main.go`. Esto permite, por ejemplo, inyectar mocks en los tests unitarios sin tocar la base de datos real.

## Cómo correr el proyecto

### Requisitos
- Docker y Docker Compose instalados. No necesitas tener Go ni PostgreSQL instalados localmente.

### Pasos

```bash
# 1. Clonar el repositorio
git clone https://github.com/BryanSLc/PROYECTO-WEB-II.git
cd PROYECTO-WEB-II

# 2. Levantar todo (API + base de datos)
docker-compose up --build
```

Con eso quedan corriendo dos contenedores:

| Servicio | Descripción | Puerto |
|---|---|---|
| `db` | PostgreSQL 16, con healthcheck | `5432:5432` |
| `app` | API en Go (Archibase) | `8080:8080` |

La API queda disponible en `http://localhost:8080`. Las tablas se crean automáticamente al arrancar (`AutoMigrate` de GORM); no se necesitan migraciones manuales.

### Variables de entorno

Configurables en el archivo `.env` (todas tienen un valor por defecto si no se define ninguna):

| Variable | Descripción | Default |
|---|---|---|
| `DB_HOST` | Host de PostgreSQL | `db` |
| `DB_PORT` | Puerto de PostgreSQL | `5432` |
| `DB_USER` | Usuario de la base de datos | `archibase` |
| `DB_PASSWORD` | Contraseña de la base de datos | `archibase` |
| `DB_NAME` | Nombre de la base de datos | `archibase` |
| `DB_SSLMODE` | Modo SSL de la conexión | `disable` |
| `APP_PORT` | Puerto donde escucha la API | `8080` |

## Endpoints por módulo

Todas las rutas están bajo el prefijo `/api/v1`. Las marcadas como **🔒 Protegida** requieren un header `Authorization: Bearer <token>` obtenido en el login.

### Auth (responsable: Bryan Lopez)

| Método | Ruta | Descripción |
|---|---|---|
| POST | `/auth/registro` | Registra un nuevo usuario (hashea la contraseña con bcrypt) |
| POST | `/auth/login` | Verifica credenciales y devuelve un JWT |

### Usuarios (responsable: Bryan Lopez)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/usuarios` | No | Crea un usuario (registro alterno) |
| GET | `/usuarios` | 🔒 | Lista todos los usuarios |
| GET | `/usuarios/{id}` | 🔒 | Obtiene un usuario por ID |
| PUT | `/usuarios/{id}` | 🔒 | Actualiza un usuario |
| DELETE | `/usuarios/{id}` | 🔒 | Elimina un usuario |

### Maquetas (responsable: Bryan Lopez)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/maquetas` | 🔒 | Crea una maqueta |
| GET | `/maquetas` | 🔒 | Lista todas las maquetas |
| GET | `/maquetas/{id}` | 🔒 | Obtiene una maqueta por ID |
| PUT | `/maquetas/{id}` | 🔒 | Actualiza una maqueta |
| DELETE | `/maquetas/{id}` | 🔒 | Elimina una maqueta |
| POST | `/maquetas/evolucion` | 🔒 | Agrega un avance/etapa a una maqueta |
| GET | `/maquetas/{id}/evolucion` | 🔒 | Lista el historial de avances de una maqueta |
| DELETE | `/maquetas/evolucion/{id}` | 🔒 | Elimina un avance |

### Recetas (responsable: Bryan Lopez)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/recetas` | 🔒 | Crea una receta (pasos de construcción) para una maqueta |
| GET | `/recetas` | 🔒 | Lista todas las recetas |
| GET | `/recetas/{id}` | 🔒 | Obtiene una receta por ID |
| PUT | `/recetas/{id}` | 🔒 | Actualiza una receta |
| DELETE | `/recetas/{id}` | 🔒 | Elimina una receta |

### Proveedores (responsable: isaac)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/proveedores` | 🔒 | Crea un proveedor de materiales |
| GET | `/proveedores` | 🔒 | Lista todos los proveedores |
| GET | `/proveedores/{id}` | 🔒 | Obtiene un proveedor por ID |
| PUT | `/proveedores/{id}` | 🔒 | Actualiza un proveedor |
| DELETE | `/proveedores/{id}` | 🔒 | Elimina un proveedor |

### Ubicaciones (responsable: isaac)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/ubicaciones` | 🔒 | Crea una ubicación (ciudad/provincia) |
| GET | `/ubicaciones` | 🔒 | Lista todas las ubicaciones |
| GET | `/ubicaciones/{id}` | 🔒 | Obtiene una ubicación por ID |
| PUT | `/ubicaciones/{id}` | 🔒 | Actualiza una ubicación |
| DELETE | `/ubicaciones/{id}` | 🔒 | Elimina una ubicación |

### Materiales (responsable: isaac)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/materiales` | 🔒 | Crea un material asociado a un proveedor |
| GET | `/materiales` | 🔒 | Lista todos los materiales |
| GET | `/materiales/{id}` | 🔒 | Obtiene un material por ID |
| PUT | `/materiales/{id}` | 🔒 | Actualiza un material |
| DELETE | `/materiales/{id}` | 🔒 | Elimina un material |

### Asesores (responsable: Eduardo García)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/asesores` | 🔒 | Crea un asesor |
| GET | `/asesores` | 🔒 | Lista todos los asesores |
| GET | `/asesores/{id}` | 🔒 | Obtiene un asesor por ID |
| PUT | `/asesores/{id}` | 🔒 | Actualiza un asesor |
| DELETE | `/asesores/{id}` | 🔒 | Elimina un asesor |

### Servicios (responsable: Eduardo García)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/servicios` | 🔒 | Crea un servicio ofrecido por un asesor |
| GET | `/servicios` | 🔒 | Lista todos los servicios |
| GET | `/servicios/{id}` | 🔒 | Obtiene un servicio por ID |
| PUT | `/servicios/{id}` | 🔒 | Actualiza un servicio |
| DELETE | `/servicios/{id}` | 🔒 | Elimina un servicio |

### Contrataciones (responsable: Eduardo García)

| Método | Ruta | Protegida | Descripción |
|---|---|---|---|
| POST | `/contrataciones` | 🔒 | Registra la contratación de un servicio |
| GET | `/contrataciones` | 🔒 | Lista todas las contrataciones |
| GET | `/contrataciones/{id}` | 🔒 | Obtiene una contratación por ID |
| PUT | `/contrataciones/{id}` | 🔒 | Actualiza una contratación |
| DELETE | `/contrataciones/{id}` | 🔒 | Elimina una contratación |

## Testing y CI

- Los tests unitarios usan mocks propios de los repositorios (implementaciones en memoria de las interfaces de cada `service`), sin depender de una base de datos real.
- El pipeline de GitHub Actions (`.github/workflows/ci.yml`) corre en cada push/PR a `main`: `go vet` → `go build` → `go test -coverprofile` contra un PostgreSQL de servicio, y sube el reporte de cobertura como artifact.

Para correr los tests localmente:

```bash
go test ./... -coverprofile=cover.out
go tool cover -func=cover.out | tail -1
```

## Diagrama Entidad-Relación

```
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
           │ 1 usuario publica muchas
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
│ tiempo_estimado       │
│ imagen               │
│ id_usuario (FK)      │
└──────────┬───────────┘
           │ 1 maqueta contiene muchos
           ▼
┌──────────────────────┐
│      MATERIAL         │
├──────────────────────┤
│ id_material (PK)     │
│ nombre               │
│ descripcion          │
│ tipo                 │
│ costo_aproximado     │
│ cantidad_aproximada  │
│ id_maqueta (FK)      │
└──────────┬───────────┘
           │ muchos materiales pueden
           │ encontrarse en muchos
           │ proveedores
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
│ tipo_material         │
└──────────────────────┘
```
