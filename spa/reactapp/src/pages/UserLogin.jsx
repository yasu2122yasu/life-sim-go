import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { useAuth } from '../AuthContext';

export const UserLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [authToken, setAuthToken] = useState('');
  const navigate = useNavigate();

  const { updateToken } = useAuth();

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/login', { email, password });
      setAuthToken(response.data.token); // 仮定：response.dataにauthTokenが含まれている
      console.log(authToken);

      if (response.data) {
        updateToken(response.data.token);
        alert('Login Success!');
        navigate('/user/user-auth');
      } else {
        alert('Login Failed!');
        navigate('/users');
      }
    } catch (error) {
      console.error('Error fetching data: ', error);
      navigate('/login');
    }
  };

  return (
    <div>
      <input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
    </div>
  );
};
