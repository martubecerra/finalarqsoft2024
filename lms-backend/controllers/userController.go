package controllers

import (
    "encoding/json"
    "net/http"
    "lms-backend/services"
    "log"
)

// Estructura para las solicitudes de registro
type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"` // Añadir campo Role
}

// Función de controlador para manejar el registro de usuarios
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error parsing request body:", err)
        http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    user, err := services.Register(req.Name, req.Email, req.Password, req.Role) // Añadir Role a la llamada del servicio
    if err != nil {
        log.Println("Error registering user:", err)
        http.Error(w, "Error registering user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    log.Printf("Registered user: %+v\n", user) // Log para verificar el usuario registrado

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// Estructura para las solicitudes de login
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Función de controlador para manejar el login
func Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error parsing request body:", err)
        http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    token, role, err := services.Authenticate(req.Email, req.Password)
    if err != nil {
        log.Println("Invalid credentials:", err)
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    log.Printf("Authenticated user: email=%s, role=%s, token=%s\n", req.Email, role, token)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
        "role":  role,
    })
}


