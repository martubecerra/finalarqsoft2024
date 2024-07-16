package main

import (
    "log"
    "net/http"
    "os"
    "time"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "lms-backend/controllers"
    "lms-backend/database"
    "lms-backend/middlewares"
    "github.com/joho/godotenv"
)

// Ruta relativa donde se guardarán los archivos dentro del contenedor Docker
const uploadPath = "./uploads"

func main() {
    // Cargar variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Agregar una pausa inicial antes de intentar conectar a la base de datos
    time.Sleep(15 * time.Second)

    // Inicializar la conexión a la base de datos
    database.Connect()

    // Configurar el archivo de log
    file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)

    // Crear un nuevo enrutador
    router := mux.NewRouter()

    // Definir las rutas y sus controladores correspondientes
    router.HandleFunc("/login", controllers.Login).Methods("POST")
    router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/courses", controllers.GetCourses).Methods("GET")
    router.HandleFunc("/courses/{id}", controllers.GetCourse).Methods("GET")

    // Rutas protegidas para usuarios (alumnos)
    router.HandleFunc("/my-courses", middlewares.AuthMiddleware(controllers.GetUserCourses)).Methods("GET")
    router.HandleFunc("/enroll", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.EnrollUser))).Methods("POST")
    router.HandleFunc("/unenroll", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.UnenrollUser))).Methods("POST")
    router.HandleFunc("/courses/{id}/comments", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.AddComment))).Methods("POST")
    router.HandleFunc("/courses/{id}/comments", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.GetComments))).Methods("GET")
    router.HandleFunc("/courses/{id}/files", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.UploadFile))).Methods("POST")
    router.HandleFunc("/courses/{id}/files", middlewares.AuthMiddleware(middlewares.UserOnly(controllers.GetFiles))).Methods("GET")

    // Rutas protegidas para administradores
    router.HandleFunc("/courses", middlewares.AuthMiddleware(controllers.CreateCourse)).Methods("POST")
    router.HandleFunc("/courses/{id}", middlewares.AuthMiddleware(controllers.UpdateCourse)).Methods("PUT")
    router.HandleFunc("/courses/{id}", middlewares.AuthMiddleware(controllers.DeleteCourse)).Methods("DELETE")

    // Servir archivos estáticos desde la carpeta uploads en la ruta relativa
    router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadPath))))

    // Configurar CORS
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}), // Cambia este origen según sea necesario
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Iniciar el servidor en el puerto 8080
    log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
