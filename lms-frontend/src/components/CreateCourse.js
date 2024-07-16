import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createCourse } from '../services/courseService';
import './CreateCourse.css'; // Asegúrate de que este archivo CSS esté en la misma carpeta

const CreateCourse = () => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [instructor, setInstructor] = useState('');
  const [duration, setDuration] = useState('');
  const [requirements, setRequirements] = useState('');
  const navigate = useNavigate();
  const token = localStorage.getItem('token');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createCourse({ title, description, instructor, duration: parseInt(duration), requirements }, token);
      navigate('/manage-courses');
    } catch (error) {
      console.error('Error creating course:', error);
    }
  };

  return (
    <div className="create-course-container">
      <h1 className="title">Create Course</h1>
      <form className="form1" onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Title:</label>
          <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Description:</label>
          <input type="text" value={description} onChange={(e) => setDescription(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Instructor:</label>
          <input type="text" value={instructor} onChange={(e) => setInstructor(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Duration:</label>
          <input type="number" value={duration} onChange={(e) => setDuration(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Requirements:</label>
          <input type="text" value={requirements} onChange={(e) => setRequirements(e.target.value)} required />
        </div>
        <button className="submit-button" type="submit">Create</button>
      </form>
    </div>
  );
};

export default CreateCourse;
