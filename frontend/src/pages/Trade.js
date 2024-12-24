// frontend/src/pages/Trade.js
import React, { useEffect, useState } from 'react';
import api from '../services/api';

function Trade() {
  const [quotes, setQuotes] = useState([]);
  const [balance, setBalance] = useState(0);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Получение котировок
        const quotesResponse = await api.get('/quotes', { params: { symbols: ['bitcoin', 'ethereum'] } });
        setQuotes(quotesResponse.data);

        // Получение баланса пользователя
        const userResponse = await api.get('/user');
        setBalance(userResponse.data.balance);
      } catch (error) {
        alert('Ошибка при получении данных');
      }
    };
    fetchData();
  }, []);

  const handleTrade = async (symbol, type) => {
    const amount = prompt(`Введите количество для ${type === 'buy' ? 'покупки' : 'продажи'}:`);
    if (amount) {
      try {
        await api.post(`/${type}`, { symbol, amount: parseFloat(amount) });
        alert(`${type === 'buy' ? 'Покупка' : 'Продажа'} успешно выполнена`);
        // Обновляем баланс
        const userResponse = await api.get('/user');
        setBalance(userResponse.data.balance);
      } catch (error) {
        alert(`Ошибка при ${type === 'buy' ? 'покупке' : 'продаже'}: ${error.response.data.error}`);
      }
    }
  };

  return (
    <div>
      <h2>Торговля</h2>
      <p>Баланс: {balance} USD</p>
      <ul>
        {quotes.map((quote) => (
          <li key={quote.symbol}>
            {quote.symbol}: {quote.price} USD
            <button onClick={() => handleTrade(quote.symbol, 'buy')}>Купить</button>
            <button onClick={() => handleTrade(quote.symbol, 'sell')}>Продать</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Trade;
