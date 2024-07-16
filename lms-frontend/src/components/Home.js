import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { getCourses } from '../services/courseService';
import './Home.css';

const Home = () => {
  const [courses, setCourses] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const data = await getCourses();
        setCourses(data);
      } catch (error) {
        console.error('Error fetching courses:', error);
      }
    };

    fetchCourses();
  }, []);

  const filteredCourses = courses.filter(course =>
    course.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="courses-container">
      <h1 className="courses-title">Courses</h1>
      <div className="search-container">
        <input
          type="text"
          className="search-input"
          placeholder="Search courses..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      <div className="courses-grid">
        {filteredCourses.map(course => (
          <div key={course.ID} className="course-card">
            <div className="card-content">
              <h2 className="course-title">{course.title}</h2>
              <Link className="course-link" to={`/courses/${course.ID}`}>
                <button className="arrow-button">
                  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                    <path d="M4 12l1.41-1.41L10 14.17V4h2v10.17l4.59-4.58L18 12l-8 8z" />
                  </svg>
                </button>
              </Link>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;
