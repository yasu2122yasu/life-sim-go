// ログインしてauthTokenが利用できるのかを確認するページ
import { Link } from 'react-router-dom';

export const AfterLogin = () => {
  return (
    <div>
      <h1>Welcome to the After Login page!</h1>
      <Link to="/user/user-auth">UserIndex</Link>
    </div>
  );
};
