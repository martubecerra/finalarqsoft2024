package models

import (
    "github.com/jinzhu/gorm"
)

// Estructura de usuario para la base de datos
type User struct {
    gorm.Model
    Name        string       `json:"name"`
    Email       string       `json:"email" gorm:"unique"`
    Password    string       `json:"-"`
    Role        string       `json:"role"`
    Enrollments []Enrollment `gorm:"foreignKey:UserID"`
    Comments    []Comment    `gorm:"foreignKey:UserID"`
    Files       []File       `gorm:"foreignKey:UserID"` // Relación con archivos
}
