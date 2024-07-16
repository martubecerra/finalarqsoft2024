import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Navigation.css';

const Navigation = ({ loggedIn, role, onLogout }) => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');
    onLogout();
    navigate('/login');
  };

  return (
    <nav className="input-horizontal">
      {loggedIn && (
        <>
          {role === 'alumno' && (
            <Link to="/home" className="value">
            Home
          </Link>
          )}
          {role === 'alumno' && (
            <Link to="/my-courses" className="value">
              My Courses
            </Link>
          )}
          {role === 'administrador' && (
            <Link to="/manage-courses" className="value">
              Manage Courses
            </Link>
          )}
          <button onClick={handleLogout} className="value">
            Logout
          </button>
        </>
      )}
    </nav>
  );
};

export default Navigation;
