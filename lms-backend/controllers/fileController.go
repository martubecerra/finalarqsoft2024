package controllers

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strconv"

    "lms-backend/database"
    "lms-backend/models"
    "lms-backend/utils"

    "github.com/gorilla/mux"
)

// UploadFile maneja la subida de archivos
func UploadFile(w http.ResponseWriter, r *http.Request) {
    userID := utils.GetUserIDFromContext(r.Context())

    vars := mux.Vars(r)
    courseID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }

    r.ParseMultipartForm(10 << 20) // 10 MB
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Crear la carpeta uploads en la raíz del proyecto si no existe
    absPath, _ := filepath.Abs("./uploads") // Ajustado para la raíz del proyecto
    if _, err := os.Stat(absPath); os.IsNotExist(err) {
        os.Mkdir(absPath, os.ModePerm)
    }

    filePath := filepath.Join(absPath, fmt.Sprintf("%d_%s", userID, handler.Filename))
    dest, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer dest.Close()

    _, err = io.Copy(dest, file)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }

    relativePath := fmt.Sprintf("uploads/%d_%s", userID, handler.Filename)
    uploadedFile := models.File{
        FileName: handler.Filename,
        FilePath: relativePath,
        UserID:   uint(userID),
        CourseID: uint(courseID),
    }

    if err := database.DB.Create(&uploadedFile).Error; err != nil {
        http.Error(w, "Error saving file record", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(uploadedFile)
}

// GetFiles obtiene todos los archivos subidos para un curso
func GetFiles(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    courseID, err := strconv.Atoi(vars["id"])
    if err != nil {
        log.Printf("Invalid course ID: %v", err)
        http.Error(w, "Invalid course ID", http.StatusBadRequest)
        return
    }
    log.Printf("CourseID: %d", courseID)

    var files []models.File
    if err := database.DB.Preload("User").Where("course_id = ?", courseID).Find(&files).Error; err != nil {
        log.Printf("Error fetching files: %v", err)
        http.Error(w, "Error fetching files", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(files)
    log.Printf("Files fetched successfully: %+v", files)
}
