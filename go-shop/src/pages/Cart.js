import React from 'react';
import { useCart } from '../context/CartContext';
import { useNavigate } from 'react-router-dom';

export const Cart = () => {
  const { cartItems, removeFromCart, updateQuantity, cartTotal } = useCart();
  const navigate = useNavigate();

  const handleCheckout = () => {
    alert("Заказ оформлен! (Здесь будет отправка на бэкенд)");
    // Очистка корзины или редирект
  };

  if (cartItems.length === 0) return <div>Корзина пуста</div>;

  return (
    <div style={{ padding: '20px' }}>
      <h1>Корзина</h1>
      {cartItems.map(item => (
        <div key={item.id} style={{ display: 'flex', justifyContent: 'space-between', borderBottom: '1px solid #ccc', padding: '10px 0' }}>
          <div>
            <h3>{item.title}</h3>
            <p>{item.price} руб.</p>
          </div>
          <div style={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
            <button onClick={() => updateQuantity(item.id, item.qty - 1)}>-</button>
            <span>{item.qty}</span>
            <button onClick={() => updateQuantity(item.id, item.qty + 1)}>+</button>
            <button onClick={() => removeFromCart(item.id)} style={{ marginLeft: '20px', color: 'red' }}>Удалить</button>
          </div>
        </div>
      ))}
      <div style={{ marginTop: '20px', textAlign: 'right' }}>
        <h2>Итого: {cartTotal} руб.</h2>
        <button onClick={handleCheckout} style={{ background: 'green', color: 'white', padding: '10px 20px', fontSize: '16px' }}>
          Оформить заказ
        </button>
      </div>
    </div>
  );
};