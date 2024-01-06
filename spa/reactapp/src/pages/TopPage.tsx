import React from 'react';
import { Link } from 'react-router-dom';

export const TopPage: React.FC = () => {
  return (
    <div className="App">
      <Link to="/">Home</Link>
      <br />
      <Link to="/user/index">UserIndex</Link>
      <br />
      <Link to="/user/login">UserLogin</Link>
      <br />
      <Link to="/user/register">UserRegister</Link>
      {/* 個別ログインユーザーのみアクセス可能 */}
      <Link to="/verify">ログイン済みユーザーのみアクセス可能</Link>
    </div>
  );
};
