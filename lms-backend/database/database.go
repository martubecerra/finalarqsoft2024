package database

import (
    "fmt"
    "log"
    "lms-backend/config"
    "lms-backend/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// Función para conectar a la base de datos
func Connect() {
    dbHost, dbPort, dbUser, dbPassword, dbName := config.GetDBConfig()
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open("mysql", connectionString)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrar las estructuras a la base de datos
    DB.AutoMigrate(&models.User{}, &models.Course{}, &models.Enrollment{}, &models.Comment{}, &models.File{})
    log.Println("Database connected and migrated")
}
