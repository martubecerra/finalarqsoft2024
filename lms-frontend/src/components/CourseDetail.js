import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { getCourseById, enrollInCourse, addComment, getComments, getFiles, uploadFile } from '../services/courseService';
import './CourseDetail.css';

const CourseDetail = () => {
  const { id } = useParams();
  const [course, setCourse] = useState(null);
  const [enrolled, setEnrolled] = useState(false);
  const [showCongrats, setShowCongrats] = useState(false);
  const [comment, setComment] = useState('');
  const [comments, setComments] = useState([]);
  const [files, setFiles] = useState([]);
  const [file, setFile] = useState(null);
  const token = localStorage.getItem('token');

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const data = await getCourseById(id);
        setCourse(data);
      } catch (error) {
        console.error('Error fetching course:', error);
      }
    };

    const fetchComments = async () => {
      try {
        const data = await getComments(id, token);
        setComments(data);
      } catch (error) {
        console.error('Error fetching comments:', error);
      }
    };

    const fetchFiles = async () => {
      try {
        const data = await getFiles(id, token);
        setFiles(data);
      } catch (error) {
        console.error('Error fetching files:', error);
      }
    };

    fetchCourse();
    fetchComments();
    fetchFiles();
  }, [id, token]);

  const handleEnroll = async () => {
    try {
      await enrollInCourse(parseInt(id), token);
      setEnrolled(true);
      setShowCongrats(true);
    } catch (error) {
      console.error('Error enrolling in course:', error);
      alert('You are already enrolled in this course.');
    }
  };

  const handleCommentSubmit = async (e) => {
    e.preventDefault();
    try {
      await addComment(id, { content: comment }, token);
      setComment('');
      const data = await getComments(id, token);
      setComments(data);
    } catch (error) {
      console.error('Error adding comment:', error);
    }
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleFileUpload = async (e) => {
    e.preventDefault();
    if (!file) return;
    try {
      const formData = new FormData();
      formData.append('file', file);
      await uploadFile(id, formData, token);
      setFile(null);
      const data = await getFiles(id, token);
      setFiles(data);
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  };

  if (!course) {
    return <div>Loading...</div>;
  }

  return (
    <div className="course-detail-container">
      <div className="course-detail-card">
        <h1 className="course-title">{course.title}</h1>
        <p className="course-description">{course.description}</p>
        <p className="course-instructor">Instructor: {course.instructor}</p>
        <p className="course-duration">Duration: {course.duration} hours</p>
        <p className="course-requirements">Requirements: {course.requirements}</p>
        {!enrolled && (
          <button className="enroll-button" onClick={handleEnroll}>
            Enroll in this course
          </button>
        )}
        {showCongrats && (
          <div className="congrats-message">
            Congrats, you have been enrolled!
          </div>
        )}
        <div className="comments-section">
          <h2>Comments</h2>
          <form onSubmit={handleCommentSubmit}>
            <textarea
              value={comment}
              onChange={(e) => setComment(e.target.value)}
              required
            ></textarea>
            <button type="submit">Add Comment</button>
          </form>
          <ul>
            {comments.map((comment) => (
              <li key={comment.id}>
                <strong>{comment.user?.name || "Unknown"}:</strong> {comment.content}
              </li>
            ))}
          </ul>
        </div>
        <div className="files-section">
          <h2>Files</h2>
          <form onSubmit={handleFileUpload}>
            <input type="file" onChange={handleFileChange} required />
            <button type="submit">Upload File</button>
          </form>
          <ul>
            {files.map((file) => (
              <li key={file.id}>
                <strong>{file.user?.name || "Unknown"}:</strong>
                <a href={`http://localhost:8080/${file.file_path.replace(/\\/g, '/')}`} download={file.file_name} target="_blank" rel="noopener noreferrer">
                  {file.file_name || "No Name"}
                </a>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default CourseDetail;
