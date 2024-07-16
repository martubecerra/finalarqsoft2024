package models

import "github.com/jinzhu/gorm"

type Comment struct {
    gorm.Model
    CourseID uint   `json:"course_id"`
    UserID   uint   `json:"user_id"`
    Content  string `json:"content"`
    User     User   `json:"user"` 
}
