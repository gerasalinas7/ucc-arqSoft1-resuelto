# Etapa de construcción: Usamos la imagen de golang para compilar la aplicación
FROM golang:1.20 AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el contenido del proyecto al contenedor
COPY . .

# Descargamos las dependencias de Go
RUN go mod tidy

# Compilamos la aplicación
RUN go build -o vinyl-store .

# Copiamos el binario compilado desde la etapa de construcción
COPY --from=builder /app/vinyl-store .

# Exponemos el puerto 8080 para la comunicación externa
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./vinyl-store"]