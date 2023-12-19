import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { loginUserAsync, AppDispatch, RootState } from './../store/index';

export const UserLogin: React.FC = () => {
  const [localEmail, setLocalEmail] = useState(''); // ローカルステートのための関数名を変更
  const [password, setPassword] = useState('');
  const dispatch = useDispatch<AppDispatch>();
  const error = useSelector((state: RootState) => state.login.error);
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      // 非同期アクション loginUserAsync を発行
      await dispatch(loginUserAsync({ email: localEmail, password }));
      navigate('/user/afterlogin');
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      <input type="text" placeholder="Email" value={localEmail} onChange={(e) => setLocalEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};
