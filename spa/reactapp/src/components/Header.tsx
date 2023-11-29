// Header.tsx
import React from 'react';

interface HeaderProps {
  title: string;
}

const Header: React.FC<HeaderProps> = ({ title }) => {
  return (
    <header style={{ backgroundColor: '#333', color: 'white', padding: '10px 20px' }}>
      <h1>{title}</h1>
    </header>
  );
};

export default Header;
