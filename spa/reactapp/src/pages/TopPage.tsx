import React from 'react';

import { Router } from './../router/Router';
import { BrowserRouter, Link } from 'react-router-dom';

export const TopPage: React.FC = () => {
  // Add your code here

  return (
    <div>
      <div className="App">
        <BrowserRouter>
          <Link to="/">Home</Link>
          <br />
          <Link to="/user/index">UserIndex</Link>
          <br />
          <Link to="/user/login">UserLogin</Link>
          <br />
          <Link to="/user/register">UserRegister</Link>
          <br />

          <Router />
        </BrowserRouter>
      </div>
    </div>
  );
};
