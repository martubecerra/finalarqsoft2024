import React, { useState, useEffect } from 'react';
import { getUserCourses, unenrollFromCourse } from '../services/courseService';
import './MyCourses.css'; // Asegúrate de tener el archivo CSS importado

const MyCourses = () => {
  const [courses, setCourses] = useState([]);
  const token = localStorage.getItem('token'); // Obtener el token del almacenamiento local

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const data = await getUserCourses(token);
        setCourses(data);
        console.log('Courses set in state:', data); // Log para verificar los cursos en el estado
      } catch (error) {
        console.error('Error fetching user courses:', error);
      }
    };

    fetchCourses();
  }, [token]);

  const handleUnenroll = async (courseId) => {
    try {
      await unenrollFromCourse(courseId, token);
      setCourses(courses.filter(course => course.ID !== courseId)); // Actualizar la lista de cursos
      console.log('Course unenrolled and removed from state:', courseId); // Log para verificar el curso removido
    } catch (error) {
      console.error('Error unenrolling from course:', error);
      alert('There was an error unenrolling from the course. Please try again.');
    }
  };

  return (
    <div className="my-courses-container">
      <h1 className="my-courses-title">My Courses</h1>
      <ul className="course-list">
        {courses.map(course => (
          <li key={course.ID} className="course-item">
            <h2 className="course-title">{course.title}</h2>
            <p></p>
            <button className="unenroll-button" onClick={() => handleUnenroll(course.ID)}>
              Unenroll
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MyCourses;
