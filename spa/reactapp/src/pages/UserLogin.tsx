// UserLogin.tsx
import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { loginUserAsync, RootState, AppDispatch } from './../store/index';

export const UserLogin: React.FC = () => {
  const [localEmail, setLocalEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const dispatch = useDispatch<AppDispatch>();
  const error = useSelector((state: RootState) => state.login.error);
  const navigate = useNavigate();

  const handleLogin = async () => {
    const actionResult = await dispatch(loginUserAsync({ email: localEmail, password }));
    if (loginUserAsync.fulfilled.match(actionResult)) {
      // ログイン成功時の処理
      navigate('/user/afterlogin');
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
