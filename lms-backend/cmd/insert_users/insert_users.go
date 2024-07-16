package main

import (
    "fmt"
    "log"
    "lms-backend/database"
    "lms-backend/models"
    "lms-backend/utils"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Inicializar la conexión a la base de datos
    database.Connect()

    // Crear usuarios de ejemplo
    users := []models.User{
        {Name: "Alice", Email: "alice@example.com", Password: hashPassword("password123"), Role: "alumno"},
        {Name: "Bob", Email: "bob@example.com", Password: hashPassword("password123"), Role: "alumno"},
        {Name: "Charlie", Email: "charlie@example.com", Password: hashPassword("password123"), Role: "administrador"},
        {Name: "David", Email: "david@example.com", Password: hashPassword("password123"), Role: "alumno"},
        {Name: "Eve", Email: "eve@example.com", Password: hashPassword("password123"), Role: "alumno"},
    }

    // Insertar usuarios en la base de datos
    for _, user := range users {
        if err := database.DB.Create(&user).Error; err != nil {
            fmt.Printf("Error creating user %s: %v\n", user.Email, err)
        } else {
            fmt.Printf("User %s created successfully\n", user.Email)
        }
    }
}

// Utilizar la función HashPassword del paquete utils
func hashPassword(password string) string {
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        log.Fatal("Error hashing password: ", err)
    }
    return hashedPassword
}
