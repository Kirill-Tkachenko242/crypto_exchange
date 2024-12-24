
import axios from 'axios';

const API_URL = "http://localhost:8080/api"; // URL бэкенда напрямую на порт 8080

const api = axios.create({
  baseURL: API_URL, // Используем прямой URL бэкенда
});

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export default api;
