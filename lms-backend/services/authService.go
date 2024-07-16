package services

import (
    "errors"
    "lms-backend/database"
    "lms-backend/models"
    "lms-backend/utils"
    "log"
)

// Función para autenticar usuarios
func Authenticate(email, password string) (string, string, error) {
    var user models.User
    database.DB.Where("email = ?", email).First(&user)

    if user.ID == 0 {
        log.Println("User not found:", email) // Log para usuario no encontrado
        return "", "", errors.New("user not found")
    }

    if !utils.CheckPasswordHash(password, user.Password) {
        log.Println("Incorrect password for user:", email) // Log para contraseña incorrecta
        return "", "", errors.New("incorrect password")
    }

    token, err := utils.GenerateJWT(user)
    if err != nil {
        log.Println("Error generating token for user:", email) // Log para error de generación de token
        return "", "", err
    }

    log.Printf("Generated token for user: email=%s, role=%s\n", email, user.Role) // Log para verificar token y rol

    return token, user.Role, nil
}

// Función para registrar usuarios
func Register(name, email, password, role string) (models.User, error) {
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return models.User{}, err
    }

    user := models.User{
        Name:     name,
        Email:    email,
        Password: hashedPassword,
        Role:     role, // Añadir campo Role
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return models.User{}, err
    }

    log.Printf("Registered new user: email=%s, role=%s\n", email, role) // Log para verificar usuario registrado

    return user, nil
}
