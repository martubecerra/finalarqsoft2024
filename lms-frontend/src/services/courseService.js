import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL;

const getCourses = async () => {
  try {
    const response = await axios.get(`${API_URL}/courses`);
    return response.data;
  } catch (error) {
    console.error('Error fetching courses:', error);
    throw error;
  }
};

const getCourseById = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/courses/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching course:', error);
    throw error;
  }
};

const createCourse = async (course, token) => {
  try {
    const response = await axios.post(`${API_URL}/courses`, course, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error creating course:', error);
    throw error;
  }
};

const updateCourse = async (id, course, token) => {
  try {
    const response = await axios.put(`${API_URL}/courses/${id}`, course, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error updating course:', error);
    throw error;
  }
};

const deleteCourse = async (id, token) => {
  try {
    await axios.delete(`${API_URL}/courses/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
  } catch (error) {
    console.error('Error deleting course:', error);
    throw error;
  }
};

const getUserCourses = async (token) => {
  try {
    const response = await axios.get(`${API_URL}/my-courses`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching user courses:', error);
    throw error;
  }
};

const enrollInCourse = async (courseId, token) => {
  try {
    const response = await axios.post(`${API_URL}/enroll`, { course_id: courseId }, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error enrolling in course:', error);
    throw error;
  }
};

const unenrollFromCourse = async (courseId, token) => {
  try {
    const response = await axios.post(`${API_URL}/unenroll`, { course_id: courseId }, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error unenrolling from course:', error);
    throw error;
  }
};

const addComment = async (courseId, commentData, token) => {
  try {
    const response = await axios.post(`${API_URL}/courses/${courseId}/comments`, commentData, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error adding comment:', error);
    throw error;
  }
};

const getComments = async (courseId, token) => {
  try {
    const response = await axios.get(`${API_URL}/courses/${courseId}/comments`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching comments:', error);
    throw error;
  }
};


const getFiles = async (courseId, token) => {
  try {
    const response = await axios.get(`${API_URL}/courses/${courseId}/files`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching files:', error);
    throw error;
  }
};


const uploadFile = async (courseId, formData, token) => {
  try {
    const response = await axios.post(`${API_URL}/courses/${courseId}/files`, formData, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data'
      }
    });
    return response.data;
  } catch (error) {
    console.error('Error uploading file:', error);
    throw error;
  }
};

export {
  getCourses,
  getCourseById,
  createCourse,
  updateCourse,
  deleteCourse,
  getUserCourses,
  enrollInCourse,
  unenrollFromCourse,
  addComment,
  getComments,
  getFiles,
  uploadFile,
};
