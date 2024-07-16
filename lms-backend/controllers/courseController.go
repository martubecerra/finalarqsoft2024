package controllers

import (
    "encoding/json"
    "net/http"
    "log"
    "lms-backend/models"
    "lms-backend/services"
    "github.com/gorilla/mux"
    "lms-backend/utils"
)

// Función de controlador para obtener todos los cursos
func GetCourses(w http.ResponseWriter, r *http.Request) {
    courses := services.GetAllCourses()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}

// Función de controlador para obtener un curso por ID
func GetCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    course := services.GetCourseByID(params["id"])
    if course.ID == 0 {
        http.Error(w, "Course not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(course)
}

// Estructura para la solicitud de creación de curso
type CreateCourseRequest struct {
    Title        string `json:"title"`
    Description  string `json:"description"`
    Instructor   string `json:"instructor"`
    Duration     int    `json:"duration"`
    UserID       uint   `json:"user_id"`
    Requirements string `json:"requirements"`
}

// Función de controlador para crear un nuevo curso
func CreateCourse(w http.ResponseWriter, r *http.Request) {
    var req CreateCourseRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error decoding course data:", err)
        http.Error(w, "Invalid course data", http.StatusBadRequest)
        return
    }
    log.Printf("Received request to create course: %+v\n", req) // Log para verificar los datos del curso

    newCourse, err := services.CreateCourse(req.Title, req.Description, req.Instructor, req.Duration, req.UserID, req.Requirements)
    if err != nil {
        log.Println("Error creating course:", err)
        http.Error(w, "Error creating course", http.StatusInternalServerError)
        return
    }

    log.Printf("Course created: %+v\n", newCourse) // Log para verificar el curso creado
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newCourse)
}

// Función de controlador para actualizar un curso existente
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var course models.Course
    err := json.NewDecoder(r.Body).Decode(&course)
    if err != nil {
        log.Println("Error decoding course data:", err)
        http.Error(w, "Invalid course data", http.StatusBadRequest)
        return
    }
    updatedCourse := services.UpdateCourse(params["id"], course)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedCourse)
}

// Función de controlador para eliminar un curso
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := services.DeleteCourse(params["id"])
    if err != nil {
        log.Println("Error deleting course:", err)
        http.Error(w, "Error deleting course", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// Función de controlador para obtener los cursos de un usuario
func GetUserCourses(w http.ResponseWriter, r *http.Request) {
    userID := utils.GetUserIDFromContext(r.Context())
    if userID == 0 {
        log.Println("Forbidden: No user ID in context")
        http.Error(w, "Forbidden: No user ID in context", http.StatusForbidden)
        return
    }

    log.Printf("Fetching courses for User ID %d\n", userID)
    courses := services.GetCoursesByUserID(userID)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}