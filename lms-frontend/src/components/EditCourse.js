import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { getCourseById, updateCourse } from '../services/courseService';

const EditCourse = () => {
  const { id } = useParams();
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [instructor, setInstructor] = useState('');
  const [duration, setDuration] = useState('');
  const [requirements, setRequirements] = useState('');
  const navigate = useNavigate();
  const token = localStorage.getItem('token');

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const data = await getCourseById(id);
        setTitle(data.title);
        setDescription(data.description);
        setInstructor(data.instructor);
        setDuration(data.duration.toString());
        setRequirements(data.requirements);
      } catch (error) {
        console.error('Error fetching course:', error);
      }
    };

    fetchCourse();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const courseData = {
        title,
        description,
        instructor,
        duration: parseInt(duration, 10),
        requirements
      };
      await updateCourse(id, courseData, token);
      navigate('/manage-courses');
    } catch (error) {
      console.error('Error updating course:', error);
    }
  };

  return (
    <div className="create-course-container">
      <h1 className="title">Edit Course</h1>
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
        <button type="submit" className="submit-button">Update</button>
      </form>
    </div>
  );
};

export default EditCourse;
