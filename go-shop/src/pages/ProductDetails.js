import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { api } from '../services/api';
import { useCart } from '../context/CartContext';

export const ProductDetails = () => {
  const { id } = useParams();
  const [product, setProduct] = useState(null);
  const { addToCart } = useCart();

  useEffect(() => {
    api.getProductById(id).then(data => setProduct(data));
  }, [id]);

  if (!product) return <div>Загрузка...</div>;

  return (
    <div style={{ display: 'flex', gap: '20px', padding: '20px' }}>
      <img src={product.image} alt={product.title} style={{ width: '300px' }} />
      <div>
        <h1>{product.title}</h1>
        <p>{product.description}</p>
        <h2>{product.price} руб.</h2>
        <p>В наличии: {product.stock} шт.</p>
        <button onClick={() => addToCart(product)} disabled={product.stock === 0}>
          {product.stock === 0 ? "Нет в наличии" : "В корзину"}
        </button>
      </div>
    </div>
  );
};