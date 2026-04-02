import React, { useEffect, useState } from 'react';
import { api } from '../services/api';
import { ProductCard } from '../components/ProductCard';

export const Home = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    // Здесь мы вызываем наш сервис
    api.getProducts().then(data => setProducts(data));
  }, []);

  return (
    <div>
      <h1>Добро пожаловать в магазин</h1>
      <p>Лучшие товары по лучшим ценам</p>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {products.map(p => <ProductCard key={p.id} product={p} />)}
      </div>
    </div>
  );
};