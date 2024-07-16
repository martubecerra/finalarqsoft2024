package services

import (
    "errors"
    "lms-backend/database"
    "lms-backend/models"
    "log"
)

// Función para obtener todos los cursos
func GetAllCourses() []models.Course {
    var courses []models.Course
    database.DB.Find(&courses)
    return courses
}

// Función para obtener un curso por ID
func GetCourseByID(id string) models.Course {
    var course models.Course
    database.DB.First(&course, id)
    return course
}

// Función para crear un nuevo curso
func CreateCourse(title, description, instructor string, duration int, userID uint, requirements string) (models.Course, error) {
    course := models.Course{
        Title:        title,
        Description:  description,
        Instructor:   instructor,
        Duration:     duration,
        UserID:       userID,
        Requirements: requirements,
    }

    if err := database.DB.Create(&course).Error; err != nil {
        log.Println("Error creating course in DB:", err)
        return models.Course{}, err
    }

    log.Printf("Created new course in DB: title=%s, instructor=%s\n", title, instructor) // Log para verificar curso creado

    return course, nil
}


// Función para actualizar un curso existente
func UpdateCourse(id string, courseData models.Course) models.Course {
    var course models.Course
    database.DB.First(&course, id)
    if course.ID == 0 {
        return models.Course{}
    }
    course.Title = courseData.Title
    course.Description = courseData.Description
    course.Instructor = courseData.Instructor
    course.Duration = courseData.Duration
    course.Requirements = courseData.Requirements
    database.DB.Save(&course)
    return course
}

// Función para eliminar un curso
func DeleteCourse(id string) error {
    var course models.Course
    database.DB.First(&course, id)
    if course.ID == 0 {
        return errors.New("course not found")
    }
    database.DB.Unscoped().Delete(&course)
    return nil
}

// Función para obtener los cursos de un usuario por su ID
func GetCoursesByUserID(userID uint) []models.Course {
    var courses []models.Course
    database.DB.Joins("JOIN enrollments ON enrollments.course_id = courses.id").Where("enrollments.user_id = ?", userID).Find(&courses)
    log.Printf("Retrieved courses for user %d: %+v\n", userID, courses)
    return courses
}
