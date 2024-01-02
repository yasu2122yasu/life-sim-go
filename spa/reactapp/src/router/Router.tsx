import React, { useState, useEffect } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { TopPage } from '../pages/TopPage';
import { UserIndex } from '../pages/UserIndex';
import { UserDetail } from '../pages/UserDetail';
import { UserLogin } from '../pages/UserLogin';
import { UserRegister } from '../pages/UserRegister';
import axios from 'axios';
import { AuthProvider, useAuth } from '../AuthContext';

export const Router: React.FC = () => {
  const authToken = useAuth();
  const [authStatus, setAuthStatus] = useState<boolean>(false);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await axios.post('http://localhost:8080/user/user-auth', { token: authToken });
        alert(response.data.authStatus);
        alert('一時停止');
        setAuthStatus(response.data.authStatus); // 仮定：レスポンスにauthStatusが含まれている
      } catch (error) {
        console.error('Error fetching data: ', error);
        setAuthStatus(false);
      }
    };

    checkAuth();
  }, [authToken]);

  return (
    <AuthProvider>
      <Routes>
        <Route path="/" element={<TopPage />} />
        <Route path="/user/index" element={<UserIndex />} />
        <Route path="/user/:userId" element={<UserDetail />} />
        <Route path="/user/login" element={<UserLogin />} />
        <Route path="/user/register" element={<UserRegister />} />
        <Route path="/user/user-auth" element={authStatus ? <UserDetail /> : <Navigate to="/user/index" />} />
      </Routes>
    </AuthProvider>
  );
};
