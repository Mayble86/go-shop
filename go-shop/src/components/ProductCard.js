import React from 'react';
import { Link } from 'react-router-dom';

export const ProductCard = ({ product }) => {
  return (
    <div style={styles.card}>
      <img src={product.image} alt={product.title} style={styles.img} />
      <h3>{product.title}</h3>
      <p>{product.price} руб.</p>
      <Link to={`/product/${product.id}`}>Подробнее</Link>
    </div>
  );
};

const styles = {
  card: { border: '1px solid #ddd', padding: '10px', margin: '10px', borderRadius: '8px', textAlign: 'center' },
  img: { width: '100%', height: '150px', objectFit: 'cover' }
};