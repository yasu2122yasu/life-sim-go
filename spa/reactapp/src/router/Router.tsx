import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { TopPage } from '../pages/TopPage';
import { UserIndex } from '../pages/UserIndex';
import { UserDetail } from '../pages/UserDetail';
import { UserLogin } from '../pages/UserLogin';
import { UserRegister } from '../pages/UserRegister';

export const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<TopPage />} />
      <Route path="/user/index" element={<UserIndex />} />
      <Route path="/user/:userId" element={<UserDetail />} />
      <Route path="/user/login" element={<UserLogin />} />
      <Route path="/user/register" element={<UserRegister />} />
      <Route path="/user/user-detail" element={<UserDetail />} />
    </Routes>
  );
};
