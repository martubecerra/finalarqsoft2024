package middlewares

import (
    "context"
    "net/http"
    "strings"
    "log"
    "lms-backend/utils"
)

// Middleware para autenticar usuarios utilizando JWT
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            log.Println("Forbidden: No token provided")
            http.Error(w, "Forbidden: No token provided", http.StatusForbidden)
            return
        }

        // Eliminar "Bearer " del token si está presente
        token = strings.TrimPrefix(token, "Bearer ")
        log.Println("Token after trim:", token)

        claims, err := utils.ParseJWT(token)
        if err != nil {
            log.Println("Forbidden: Invalid token", err)
            http.Error(w, "Forbidden: Invalid token", http.StatusForbidden)
            return
        }

        log.Printf("Token valid: UserID=%d, Role=%s\n", claims.UserID, claims.Role)

        // Añadir los claims al contexto de la petición
        ctx := context.WithValue(r.Context(), utils.ClaimsKey, claims)
        r = r.WithContext(ctx)

        next(w, r)
    })
}
