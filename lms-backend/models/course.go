// models/course.go
package models

import (
    "github.com/jinzhu/gorm"
)

// Estructura de curso para la base de datos
type Course struct {
    gorm.Model
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Instructor  string     `json:"instructor"`
    Duration    int        `json:"duration"` // duración en horas
    UserID      uint       `json:"user_id"`
    Requirements string    `json:"requirements"` // nuevos requisitos
    Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
    Comments    []Comment  `gorm:"foreignKey:CourseID"` // Añadir relación con Comment
    Files       []File     `gorm:"foreignKey:CourseID"` // Añadir relación con File
}
