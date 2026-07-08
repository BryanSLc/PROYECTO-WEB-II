# ==========================================
# ETAPA 1: Compilación
# ==========================================
FROM golang:1.26-alpine AS build

WORKDIR /app

# Dependencias primero (mejor uso de caché de capas)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos un binario estático (CGO desactivado: ya no necesitamos SQLite/cgo)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/archibase ./cmd/api

# ==========================================
# ETAPA 2: Imagen final (liviana)
# ==========================================
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

# Usuario no-root por seguridad: la app no corre como root dentro del contenedor
RUN adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=build /app/archibase .

USER appuser

EXPOSE 8080

CMD ["./archibase"]