package models

import (
    "time"
)

type Enrollment struct {
    ID        uint      `gorm:"primary_key"`
    UserID    uint      `json:"user_id" gorm:"unique_index:idx_user_course"`
    CourseID  uint      `json:"course_id" gorm:"unique_index:idx_user_course"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
