# Usa una imagen base de Node.js
FROM node:14

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Instala las dependencias del proyecto
RUN npm install

# Compila la aplicación
RUN npm run build

# Sirve la aplicación usando `serve`
RUN npm install -g serve

# Expone el puerto en el que correrá la aplicación
EXPOSE 3000

# Ejecuta la aplicación
CMD ["serve", "-s", "build"]
