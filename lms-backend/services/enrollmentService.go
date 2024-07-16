package services

import (
    "lms-backend/database"
    "lms-backend/models"
    "errors"
)

// Función para inscribir a un usuario en un curso
func EnrollUser(userID, courseID uint) (models.Enrollment, error) {
    var existingEnrollment models.Enrollment
    if database.DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existingEnrollment).RecordNotFound() {
        enrollment := models.Enrollment{
            UserID:   userID,
            CourseID: courseID,
        }
        database.DB.Create(&enrollment)
        return enrollment, nil
    }
    return models.Enrollment{}, errors.New("user already enrolled in this course")
}

// Función para obtener los cursos de un usuario
func GetUserCourses(userID uint) []models.Course {
    var courses []models.Course
    database.DB.Table("courses").
        Select("courses.*").
        Joins("join enrollments on enrollments.course_id = courses.id").
        Where("enrollments.user_id = ?", userID).
        Scan(&courses)
    return courses
}

// Función para desinscribir a un usuario de un curso
func UnenrollUser(userID, courseID uint) error {
    if err := database.DB.Unscoped().Where("user_id = ? AND course_id = ?", userID, courseID).Delete(&models.Enrollment{}).Error; err != nil {
        return err
    }
    return nil
}
