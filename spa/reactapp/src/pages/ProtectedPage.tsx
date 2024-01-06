import React from 'react';
import { Link } from 'react-router-dom';

const PublicPage: React.FC = () => {
  // CSSを記述する
  const textStyle = {
    color: 'blue',
    fontWeight: 'bold',
  };

  return (
    <div>
      <h1>Public Page</h1>
      <p>This page can be accessed by any logged-in user.</p>
      <Link to="/" style={textStyle}>
        Homeに戻る
      </Link>
    </div>
  );
};

export default PublicPage;
