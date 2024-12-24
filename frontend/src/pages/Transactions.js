// frontend/src/pages/Transactions.js
import React, { useEffect, useState } from 'react';
import api from '../services/api';

function Transactions() {
  const [transactions, setTransactions] = useState([]);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await api.get('/transactions');
        setTransactions(response.data);
      } catch (error) {
        alert('Ошибка при получении транзакций');
      }
    };
    fetchTransactions();
  }, []);

  return (
    <div>
      <h2>История транзакций</h2>
      <table>
        <thead>
          <tr>
            <th>Тип</th>
            <th>Сумма</th>
            <th>Дата</th>
          </tr>
        </thead>
        <tbody>
          {transactions.map((tx) => (
            <tr key={tx.id}>
              <td>{tx.type}</td>
              <td>{tx.amount} USD</td>
              <td>{new Date(tx.created_at).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Transactions;
