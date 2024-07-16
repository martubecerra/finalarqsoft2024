package models

import (
    "github.com/jinzhu/gorm"
)

// File estructura para representar un archivo subido
type File struct {
    gorm.Model
    FileName string `json:"file_name"`
    FilePath string `json:"file_path"`
    UserID   uint   `json:"user_id"`
    CourseID uint   `json:"course_id"`
    User     User   `json:"user"` 
}
