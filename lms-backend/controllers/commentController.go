package controllers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "lms-backend/database"
    "lms-backend/models"
    "lms-backend/utils"
    "github.com/gorilla/mux"
)

// AddComment añade un comentario a un curso
func AddComment(w http.ResponseWriter, r *http.Request) {
    userID := utils.GetUserIDFromContext(r.Context())
    log.Println("Received request to add comment for user ID:", userID)

    // Obtener el ID del curso de la URL
    vars := mux.Vars(r)
    courseID, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Println("Error parsing course ID:", err)
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }
    log.Println("Parsed course ID:", courseID)

    // Decodificar el cuerpo de la solicitud
    var comment models.Comment
    err = json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        log.Println("Error decoding comment:", err)
        http.Error(w, "Error decoding comment", http.StatusBadRequest)
        return
    }

    // Establecer el CourseID y UserID
    comment.CourseID = uint(courseID)
    comment.UserID = userID
    log.Println("Decoded comment:", comment)

    // Guardar el comentario en la base de datos
    if err := database.DB.Create(&comment).Error; err != nil {
        log.Println("Error saving comment:", err)
        http.Error(w, "Error saving comment", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(comment)
}

// GetComments obtiene todos los comentarios de un curso
func GetComments(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    courseID, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Println("Error parsing course ID:", err)
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }

    var comments []models.Comment
    if err := database.DB.Preload("User").Where("course_id = ?", courseID).Find(&comments).Error; err != nil {
        log.Println("Error fetching comments:", err)
        http.Error(w, "Error fetching comments", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comments)
}
