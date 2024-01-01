import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { checkLoginUser } from '../features/login/LoginSlice';
import axios from 'axios';

export const UserLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [authToken, setAuthToken] = useState('');
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/login', { email, password });
      setAuthToken(response.data.token); // 仮定：response.dataにauthTokenが含まれている

      console.log(authToken);
      if (response.data) {
        dispatch(checkLoginUser({ email, authToken: response.data }));
        navigate('/user/user-detail');
      } else {
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
