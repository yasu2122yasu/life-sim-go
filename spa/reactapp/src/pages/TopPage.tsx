import React from 'react';
import { Link } from 'react-router-dom';
import { Router } from './../router/Router';

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
      <br />
    </div>
  );
};
