import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Header from './Header';
import SideBar from './SideBar';

interface UserData {
  ID: number;
  Name: string;
  Email: string;
  Password: string;
  CreatedAt: string;
  UpdatedAt: string;
}

function UserIndex() {
  const [userData, setUserData] = useState<UserData[]>([]); // ユーザーデータの配列

  useEffect(() => {
    axios
      .get('http://localhost:8080/users')
      .then((response) => {
        setUserData(response.data); // レスポンスデータを配列として設定
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, []); // コンポーネントがマウントされたときに一度だけ実行

  if (userData.length === 0) return <div>Loading...</div>;

  return (
    <div>
      <Header title="User Information" />
      <div style={{ float: 'right' }} className="sidebar">
        <SideBar />
      </div>
      <h1>User Data</h1>
      {userData.map((user) => (
        <div key={user.ID}>
          <p>ID: {user.ID}</p>
          <p>Name: {user.Name}</p>
          <p>Email: {user.Email}</p>
          <p>Created At: {user.CreatedAt}</p>
          <p>Updated At: {user.UpdatedAt}</p>
        </div>
      ))}
    </div>
  );
}

export default UserIndex;
