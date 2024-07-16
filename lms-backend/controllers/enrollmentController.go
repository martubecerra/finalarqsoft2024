package controllers

import (
    "encoding/json"
    "net/http"
    "log"
    "lms-backend/services"
    "lms-backend/utils"
)

// Estructura para la solicitud de inscripción y desinscripción
type EnrollRequest struct {
    CourseID uint `json:"course_id"`
}

// Función de controlador para inscribir a un usuario en un curso
func EnrollUser(w http.ResponseWriter, r *http.Request) {
    var req EnrollRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error parsing request body:", err)
        http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    userID := utils.GetUserIDFromContext(r.Context())
    if userID == 0 {
        log.Println("Forbidden: No user ID in context")
        http.Error(w, "Forbidden: No user ID in context", http.StatusForbidden)
        return
    }

    log.Printf("User ID %d enrolling in course %d\n", userID, req.CourseID)

    enrollment, err := services.EnrollUser(userID, req.CourseID)
    if err != nil {
        log.Println("Error enrolling user:", err)
        http.Error(w, "Error enrolling user: "+err.Error(), http.StatusConflict)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(enrollment)
}

// Función de controlador para desinscribir a un usuario de un curso
func UnenrollUser(w http.ResponseWriter, r *http.Request) {
    var req EnrollRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error parsing request body:", err)
        http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    userID := utils.GetUserIDFromContext(r.Context())
    if userID == 0 {
        log.Println("Forbidden: No user ID in context")
        http.Error(w, "Forbidden: No user ID in context", http.StatusForbidden)
        return
    }

    log.Printf("User ID %d unenrolling from course %d\n", userID, req.CourseID)

    err = services.UnenrollUser(userID, req.CourseID)
    if err != nil {
        log.Println("Error unenrolling user:", err)
        http.Error(w, "Error unenrolling user: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
