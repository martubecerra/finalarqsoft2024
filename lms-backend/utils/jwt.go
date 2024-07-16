package utils

import (
    "context"
    "time"
    "lms-backend/models"
    "github.com/dgrijalva/jwt-go"
    "log"
)

var jwtKey = []byte("3a%J%wXVb0!W&FCbK*'5") // Define la clave secreta aquí

// Definir un tipo de clave personalizado
type ContextKey string

const (
    ClaimsKey ContextKey = "claims"
)

// Estructura para el token JWT
type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

// Función para generar un token JWT
func GenerateJWT(user models.User) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: user.ID,
        Role:   user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// Función para parsear un token JWT
func ParseJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        log.Println("Error parsing token:", err)
        return nil, err
    }

    if !token.Valid {
        log.Println("Password or email incorrect")
        return nil, err
    }

    return claims, nil
}

// Función para obtener el ID del usuario desde el contexto
func GetUserIDFromContext(ctx context.Context) uint {
    claims, _ := ctx.Value(ClaimsKey).(*Claims)
    return claims.UserID
}
