import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

export const UserRegister = () => {
  // CSSを記述する
  const textStyle = {
    color: 'blue',
    fontWeight: 'bold',
  };

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleRegister = async () => {
    try {
      const response = await axios.post('http://localhost:8080/users', { email, password });
      console.log(response);
      navigate('/');
    } catch (error) {
      console.error('Error fetching data: ', error);
      navigate('/login');
    }
  };

  return (
    <div>
      <input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleRegister}>ユーザーを登録!</button>
      <Link to="/" style={textStyle}>
        Homeに戻る
      </Link>
    </div>
  );
};
