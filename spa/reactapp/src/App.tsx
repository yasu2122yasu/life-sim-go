import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface UserData {
  ID: number;
  Name: string;
  Email: string;
  Password: string;
  CreatedAt: string;
  UpdatedAt: string;
}

function App() {
  const [userData, setUserData] = useState<UserData | null>(null);

  useEffect(() => {
    axios
      .get('http://localhost:8080/users/1')
      .then((response) => {
        setUserData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, []);

  if (!userData) return <div>Loading...</div>;

  return (
    <div>
      <h1>User Data</h1>
      <p>ID: {userData.ID}</p>
      <p>Name: {userData.Name}</p>
      <p>Email: {userData.Email}</p>
      {/* パスワードは通常表示しない */}
      <p>Created At: {userData.CreatedAt}</p>
      <p>Updated At: {userData.UpdatedAt}</p>
    </div>
  );
}

export default App;
