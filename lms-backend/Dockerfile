# Usa una imagen base de Golang
FROM golang:1.18

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum primero para aprovechar el caché de Docker
COPY go.mod go.sum ./

# Descarga las dependencias del módulo
RUN go mod download

# Copia el resto de los archivos del proyecto
COPY . .

# Establece el directorio de trabajo donde se encuentra el main.go
WORKDIR /app/cmd

# Compila la aplicación
RUN go build -o /app/main .

# Establece el directorio de trabajo de nuevo a /app
WORKDIR /app

# Expone el puerto en el que correrá la aplicación
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./main"]