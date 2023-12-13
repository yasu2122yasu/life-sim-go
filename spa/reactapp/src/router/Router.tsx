import { Routes, Route, BrowserRouter, Link } from 'react-router-dom';
import { TopPage } from '../pages/TopPage';
import { UserIndex } from '../pages/UserIndex';
import { UserLogin } from '../pages/UserLogin';
import { UserRegister } from '../pages/UserRegister';

export const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<TopPage />} />
      <Route path="/user/index" element={<UserIndex />} />
      <Route path="/user/login" element={<UserLogin />} />
      <Route path="/user/register" element={<UserRegister />} />
    </Routes>
  );
};
