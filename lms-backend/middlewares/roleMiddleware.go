package middlewares

import (
    "net/http"
    "lms-backend/utils"
    "log"
)

// Middleware para verificar el rol del usuario
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims := r.Context().Value(utils.ClaimsKey).(*utils.Claims)
        if claims.Role != "administrador" {
            log.Println("Forbidden: Not an administrator")
            http.Error(w, "Forbidden: Not an administrator", http.StatusForbidden)
            return
        }
        next(w, r)
    })
}

func UserOnly(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        claims := r.Context().Value(utils.ClaimsKey).(*utils.Claims)
        if claims.Role != "alumno" {
            log.Println("Forbidden: Not a user")
            http.Error(w, "Forbidden: Not a user", http.StatusForbidden)
            return
        }
        next(w, r)
    })
}
