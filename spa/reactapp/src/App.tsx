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
  const [userId, setUserId] = useState<number>(1); // IDの初期値を設定
  const [userData, setUserData] = useState<UserData | null>(null);

  useEffect(() => {
    axios
      .get(`http://localhost:8080/users/${userId}`) // テンプレートリテラルを使用してIDを動的に設定
      .then((response) => {
        setUserData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, [userId]); // userIdが変更されたときにuseEffectを再実行

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
      {/* ユーザーIDを変更するための入力フィールドとボタン */}
      <input type="number" value={userId} onChange={(e) => setUserId(parseInt(e.target.value, 10))} />
    </div>
  );
}

export default App;
