import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getCourses, deleteCourse } from '../services/courseService';
import './ManageCourses.css'; // Asegúrate de tener el archivo CSS importado

const ManageCourses = () => {
  const [courses, setCourses] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCourses = async () => {
      const data = await getCourses();
      setCourses(data);
    };

    fetchCourses();
  }, []);

  const handleDelete = async (courseId) => {
    if (window.confirm('Are you sure you want to delete this course?')) {
      await deleteCourse(courseId, localStorage.getItem('token'));
      setCourses(courses.filter(course => course.ID !== courseId));
    }
  };

  const handleAction = (action, courseId = null) => {
    if (action === 'create') {
      navigate('/create-course');
    } else if (action === 'edit' && courseId) {
      navigate(`/edit-course/${courseId}`);
    }
  };

  return (
    <div className="manage-courses-container">
      <h1 className="manage-courses-title">Manage Courses</h1>
      <button className="create-course-button" onClick={() => handleAction('create')}>Create Course</button>
      <ul className="course-list">
        {courses.map(course => (
          <li key={course.ID} className="course-item">
            <span className="course-title">{course.title}</span>
            <div className="course-actions">
              <button className="edit-button" onClick={() => handleAction('edit', course.ID)}>Edit</button>
              <button className="delete-button" onClick={() => handleDelete(course.ID)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ManageCourses;
