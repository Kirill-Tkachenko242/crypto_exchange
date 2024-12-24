// frontend/src/components/NavBar.js
import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

function NavBar() {
  const navigate = useNavigate();
  const isAuthenticated = !!localStorage.getItem('token');

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <nav>
      {isAuthenticated ? (
        <>
          <Link to="/">Торговля</Link>
          <Link to="/transactions">Транзакции</Link>
          <button onClick={handleLogout}>Выйти</button>
        </>
      ) : (
        <>
          <Link to="/login">Войти</Link>
          <Link to="/register">Регистрация</Link>
        </>
      )}
    </nav>
  );
}

export default NavBar;
