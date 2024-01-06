import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../AuthContext';

export const UserLogin = () => {
  // CSSを記述する
  const textStyle = {
    color: 'blue',
    fontWeight: 'bold',
  };

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [authToken, setAuthToken] = useState('');
  const navigate = useNavigate();

  const { updateToken } = useAuth();

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/login', { email, password });
      console.log(response.data);
      setAuthToken(response.data.token); // 仮定：response.dataにauthTokenが含まれている

      if (response.data) {
        updateToken(response.data.token);
        alert('Login Success!');
        navigate('/verify');
      } else {
        alert('Login Failed!');
        navigate('/users');
      }
    } catch (error) {
      console.error('Error fetching data: ', error);
      //   navigate('/login');
    }
  };

  return (
    <div>
      <input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
      <Link to="/" style={textStyle}>
        Homeに戻る
      </Link>
    </div>
  );
};
