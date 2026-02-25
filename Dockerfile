# --- Etapa 1: Compilación ---
# Usamos la versión exacta de tu go.mod
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copiamos los archivos de dependencias y descargamos
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el código fuente
COPY . .

# Compilamos la aplicación apuntando específicamente a tu main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sharedup-api ./app/main.go

# --- Etapa 2: Producción ---
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiamos SOLO el binario compilado de la etapa anterior
COPY --from=builder /app/sharedup-api .

# Exponemos el puerto 8081 que usas en tu main.go
EXPOSE 8081

# Ejecutamos la app
CMD ["./sharedup-api"]