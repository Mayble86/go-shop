import React from 'react';
import { Link } from 'react-router-dom';
import { useCart } from '../context/CartContext';
import { useAuth } from '../context/AuthContext';

export const Header = () => {
  const { cartItems } = useCart();
  const { user, logout } = useAuth();

  return (
    <header style={styles.header}>
      <div style={styles.logo}><Link to="/">MyShop</Link></div>
      <nav>
        <Link to="/cart">Корзина ({cartItems.length})</Link>
        {user ? (
          <>
            <span>Привет, {user.name}</span>
            <button onClick={logout}>Выйти</button>
          </>
        ) : (
          <Link to="/login">Войти</Link>
        )}
      </nav>
    </header>
  );
};

const styles = {
  header: { display: 'flex', justifyContent: 'space-between', padding: '20px', background: '#333', color: '#fff' },
  logo: { fontSize: '24px', fontWeight: 'bold' }
};