# finalarqsoft2024

Proyecto Práctico Integrador 2024

# Arquitectura de Software I - 2024

## Integrantes del Grupo

- *Martina Becerra* (2214822)
- *Sofía Contreras* (2215803)
- *Manuel Frías Faoro* (2217890)
- *José Ruarte* (2206224)


## Enunciado Práctico Integrador 2024

Como práctico integrador se solicita la creación de un sistema de gestión de cursos en línea (LMS), donde se destacan dos componentes a ser desarrollados:

- *Backend*, desarrollado en Golang, brindará todas las interfaces necesarias para dar solución al requerimiento.
- *Frontend*, desarrollado en React, representará la vista final del usuario y consumirá los servicios desarrollados en el backend.

## Puntos Requeridos para la Construcción del Sistema

### Backend
- *Autenticación de usuarios*: Implementar un sistema de login y gestión de permisos de usuarios. Deben existir 2 tipos de usuarios: alumno y administrador.
- *Gestión de cursos*: Desarrollar endpoints que permitan la creación, edición, y eliminación de cursos por parte de los administradores.
- *Gestión de usuarios inscritos*: Implementar un endpoint para listar los cursos a los que un usuario está inscrito.
- *Seguridad*: Garantizar la seguridad de las transacciones (autorización por token firmado entre frontend y backend) y datos (hashing de contraseñas) del sistema.

### Frontend
- *Pantalla de inicio (Home)*: Mostrar un listado de cursos disponibles para inscripción.
- *Búsqueda de cursos*: Implementar un motor de búsqueda que permita a los usuarios encontrar cursos por palabras clave o categorías.
- *Detalle del curso*: Mostrar información detallada sobre un curso seleccionado, incluyendo descripción, instructor, duración, y requisitos.
- *Inscripción en curso*: Habilitar un botón de inscripción para que los usuarios puedan registrarse en los cursos de su interés.
- *Mis cursos*: Mostrar un listado de los cursos a los que el usuario está inscrito, con la opción de acceder a los detalles de cada curso.

## Condiciones de Regularidad y Examen Final

### Regularidad
Para regularizar la materia se pide el desarrollo relacionado con:
- Autenticación de usuarios.
- Visualización de la página de inicio con el listado de cursos.
- Búsqueda de cursos.
- Funcionalidad de inscripción.

### Examen Final
Para el examen final se solicita el desarrollo completo del sistema y puntos extras enumerados a continuación:
- Gestión de cursos por parte de los administradores.
- Seguridad del sistema en todos los componentes.
- Funcionalidad de listado de cursos a los que está inscrito el usuario.
- Implementar un sistema de comentarios y valoraciones para los cursos.
- Permitir a los usuarios subir archivos relacionados con los cursos.
- Dockerizar y componer la solución completa.

## Criterios de Evaluación - Primera Entrega Parcial - Final

### Frontend
- Se muestra correctamente el formulario de login para usuario alumno y administrador.
- Se muestra correctamente la home con listado/búsqueda de cursos para seleccionar.
- Se muestra correctamente el detalle del curso seleccionado.
- Se muestra correctamente la página o mensaje de congrats al inscribirse a un curso.
- Se muestra correctamente la página de mis cursos para el usuario alumno.
- Se muestra correctamente el formulario/resultado de agregar un comentario.
- Permite al usuario subir archivos relacionados a los cursos.
- Se muestra correctamente la carga de un curso para el usuario administrador.

### Backend
- Implementa login recibiendo las credenciales y genera el token de acceso JWT.
- Implementa correctamente la búsqueda de cursos.
- Respeta estructura MVC y responsabilidades de modelo, vista y controlador y DTO.
- No ignora errores y retorna correctamente los códigos de estado.
- Implementa correctamente la creación de la curso.
- Implementa correctamente el listado de cursos para un usuario.
- Implementa correctamente el agregado y listado de comentarios.
- Permite al usuario subir archivos relacionados a los cursos.

### BBDD
- La base de datos no almacena la información de la password de forma plana.
- Se conecta correctamente contra la base de datos usando GORM.
- Las distintas tablas en la base no duplican la información en más de una tabla.

### General
- El código fuente del proyecto está subido a Github.
- La aplicación completa se puede correr utilizando Docker.


## Pasos para Ejecutar el Programa

1. Crear una base de datos MySQL, con la siguiente consulta:
    sql
    -- Crear la base de datos
    CREATE DATABASE lms_db;

    -- Usar la base de datos
    USE lms_db;

    -- Crear la tabla users
    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE,
        password VARCHAR(255),
        role ENUM('alumno', 'administrador'),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

    -- Crear la tabla courses
    CREATE TABLE courses (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(100),
        description TEXT,
        instructor VARCHAR(100),
        duration INT,
        requirements TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

    -- Insertar datos en la tabla courses
    INSERT INTO courses (title, description, instructor, duration, requirements) VALUES
    ('Curso de Desarrollo Web Full Stack con Golang y React', 
     'Este curso intensivo de 12 semanas te llevará desde los conceptos básicos hasta el dominio avanzado del desarrollo web full stack. Aprenderás a construir aplicaciones web completas utilizando Golang en el backend y React en el frontend. Incluye módulos sobre autenticación, gestión de bases de datos, y despliegue en servidores en la nube.',
     'Laura Martínez', 120, 'Conocimientos básicos de programación.'),

    ('Masterclass en Diseño de Interfaces de Usuario (UI) y Experiencia de Usuario (UX)', 
     'En esta masterclass, aprenderás a diseñar interfaces de usuario efectivas y atractivas, así como a mejorar la experiencia de usuario en aplicaciones web y móviles. El curso cubre principios de diseño visual, prototipado, y pruebas de usabilidad. También se incluyen estudios de caso de empresas líderes en la industria.',
     'Fernando García', 80, 'Conocimientos básicos de diseño y usabilidad.'),

    ('Curso de Ciencia de Datos y Machine Learning con Python', 
     'Este curso avanzado está diseñado para aquellos interesados en la ciencia de datos y el machine learning. A lo largo de 10 semanas, explorarás técnicas de análisis de datos, visualización, y algoritmos de machine learning. Utilizarás Python y librerías populares como Pandas, NumPy, y Scikit-Learn para resolver problemas complejos y obtener insights valiosos de grandes volúmenes de datos.',
     'Ana Rodríguez', 100, 'Experiencia previa en programación y análisis de datos.'),

    ('Curso de Programación en Java', 
     'Este curso introductorio te enseñará los fundamentos de la programación en Java. Aprenderás sobre variables, estructuras de control, funciones y clases, sentando las bases para tu carrera como desarrollador Java.',
     'Pedro Sánchez', 60, 'No se requieren conocimientos previos.'),

    ('Taller de Desarrollo de Aplicaciones Móviles con Flutter', 
     'Descubre cómo desarrollar aplicaciones móviles multiplataforma con Flutter. En este taller práctico, aprenderás a crear interfaces de usuario hermosas y fluidas, gestionar estados, y conectar tu aplicación con servicios en la nube.',
     'María López', 40, 'Conocimientos básicos de programación.');

    -- Crear la tabla enrollments
    CREATE TABLE enrollments (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        course_id INT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        UNIQUE KEY idx_user_course (user_id, course_id),
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (course_id) REFERENCES courses(id)
    );

    -- Crear la tabla comments
    CREATE TABLE comments (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    course_id INT(10) UNSIGNED DEFAULT NULL,
    user_id INT(10) UNSIGNED DEFAULT NULL,
    content VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id),
    KEY deleted_at (deleted_at)
    );

    -- Crear la tabla files
    CREATE TABLE files (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    file_name VARCHAR(255) DEFAULT NULL,
    file_path VARCHAR(255) DEFAULT NULL,
    user_id INT(10) UNSIGNED DEFAULT NULL,
    course_id INT(10) UNSIGNED DEFAULT NULL,
    PRIMARY KEY (id),
    KEY deleted_at (deleted_at)
    );
    

2. Ejecutar el backend desde la terminal:
    bash
    cd lms-backend
    go run cmd/main.go
    

3. Ejecutar el frontend desde la terminal:
    bash
    cd lms-frontend
    npm start
    
4. Ejecutar desde docker, abrir una terminal en la carpeta del archivo:
    bash
    docker-compose down
    docker-compose up --build
    