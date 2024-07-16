import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Home from './components/Home';
import Login from './components/Login';
import CourseDetail from './components/CourseDetail';
import MyCourses from './components/MyCourses';
import ManageCourses from './components/ManageCourses';
import CreateCourse from './components/CreateCourse';
import EditCourse from './components/EditCourse';
import Navigation from './components/Navigation';
import Register from './components/Register';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [role, setRole] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      try {
        const decodedToken = JSON.parse(atob(token.split('.')[1]));
        setLoggedIn(true);
        setRole(decodedToken.role);
      } catch (error) {
        console.error('Error decoding token:', error);
        localStorage.removeItem('token');
      }
    }
  }, []);

  const handleLogin = () => {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = JSON.parse(atob(token.split('.')[1]));
      setLoggedIn(true);
      setRole(decodedToken.role);
    }
  };

  const handleLogout = () => {
    setLoggedIn(false);
    setRole(null);
    localStorage.removeItem('token');
  };

  return (
    <Router>
      <div>
        <Navigation loggedIn={loggedIn} role={role} onLogout={handleLogout} />
        <Routes>
          <Route path="/login" element={<Login onLogin={handleLogin} />} />
          <Route path="/register" element={<Register />} />
          <Route path="/home" element={loggedIn ? <Home /> : <Navigate to="/login" />} />
          <Route path="/courses/:id" element={loggedIn ? <CourseDetail /> : <Navigate to="/login" />} />
          <Route path="/my-courses" element={loggedIn && role === 'alumno' ? <MyCourses /> : <Navigate to="/login" />} />
          <Route path="/manage-courses" element={loggedIn && role === 'administrador' ? <ManageCourses /> : <Navigate to="/login" />} />
          <Route path="/create-course" element={loggedIn && role === 'administrador' ? <CreateCourse /> : <Navigate to="/login" />} />
          <Route path="/edit-course/:id" element={loggedIn && role === 'administrador' ? <EditCourse /> : <Navigate to="/login" />} />
          <Route path="/" element={<Navigate to="/login" />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
