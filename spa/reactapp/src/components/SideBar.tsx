import React, { useState } from 'react';
import UserIndex from './UserIndex';

const SideBar: React.FC = () => {
  const [showUserIndex, setShowUserIndex] = useState(false);

  const handleClick = () => {
    setShowUserIndex(true);
  };

  return (
    <aside style={{ backgroundColor: '#f2f2f2', width: '200px', padding: '10px' }}>
      <button onClick={handleClick}>Click me</button>
      {showUserIndex && <UserIndex />}
    </aside>
  );
};

export default SideBar;
