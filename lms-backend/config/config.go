package config

import (
    "os"
)

// Función para obtener la configuración de la base de datos desde las variables de entorno
func GetDBConfig() (string, string, string, string, string) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    return dbHost, dbPort, dbUser, dbPassword, dbName
}
