import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { loginUser } from './../auth/authSlice';
import { RootState, AppDispatch } from './../auth/store';

export const UserLogin: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const dispatch = useDispatch<AppDispatch>();
  const error = useSelector((state: RootState) => state.auth.error);
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      dispatch(loginUser({ email, password })).unwrap();
      navigate('/user/afterlogin'); // ログイン成功時にリダイレクト
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      <input type="text" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};
