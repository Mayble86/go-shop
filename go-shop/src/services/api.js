// Имитация базы данных товаров
const MOCK_PRODUCTS = [
  { id: 1, title: "Товар 1", price: 1000, description: "Описание товара 1", image: "https://via.placeholder.com/300", stock: 10 },
  { id: 2, title: "Товар 2", price: 2500, description: "Описание товара 2", image: "https://via.placeholder.com/300", stock: 5 },
  { id: 3, title: "Товар 3", price: 500, description: "Описание товара 3", image: "https://via.placeholder.com/300", stock: 0 },
];

// Имитация задержки сети
const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms));

export const api = {
  // Получить все товары
  getProducts: async () => {
    await delay(500); 
    return MOCK_PRODUCTS;
  },

  // Получить один товар
  getProductById: async (id) => {
    await delay(500);
    return MOCK_PRODUCTS.find(p => p.id === parseInt(id));
  },

  // Имитация логина
  login: async (email, password) => {
    await delay(500);
    if (email && password) return { id: 1, name: "User", email };
    throw new Error("Неверные данные");
  }
};