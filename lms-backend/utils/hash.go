package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// Función para generar el hash de una contraseña
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Función para verificar una contraseña con su hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
