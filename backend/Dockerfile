# Usamos la imagen oficial de Go que coincide con tu go.mod
FROM golang:1.24.3

# Creamos y seleccionamos el directorio de trabajo
WORKDIR /app

# Primero copiamos solo los archivos de módulo (para mejor caching)
COPY go.mod go.sum ./

# Descargamos las dependencias
RUN go mod download

# Copiamos todo el código fuente
COPY . .

# Compilamos la aplicación
RUN go build -o main .

# Exponemos el puerto que usa tu aplicación
EXPOSE 8000

# Comando para ejecutar la aplicación
CMD ["./main"]