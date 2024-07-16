import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL;

export const login = async (email, password) => {
  try {
    const response = await axios.post(`${API_URL}/login`, { email, password });
    const { token, role } = response.data;
    console.log('Login response:', response.data); // Verificar la respuesta
    localStorage.setItem('token', token);
    localStorage.setItem('role', role); // Almacenar el rol
    return { token, role };
  } catch (error) {
    console.error('Error during login:', error);
    throw error;
  }
};

export const register = async (name, email, password, role) => {
  try {
    const response = await axios.post(`${API_URL}/register`, { name, email, password, role });
    return response.data;
  } catch (error) {
    console.error('Error during registration:', error);
    throw error;
  }
};
