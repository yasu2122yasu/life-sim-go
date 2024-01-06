// ログインしてauthTokenが利用できるのかを確認するページ
import { Link } from 'react-router-dom';

export const AfterLogin = () => {
  // CSSを記述する
  const textStyle = {
    color: 'blue',
    fontWeight: 'bold',
  };
  return (
    <div>
      <h1>Welcome to the After Login page!</h1>
      <Link to="/user/user-auth">UserIndex</Link>
      <Link to="/" style={textStyle}>
        Homeに戻る
      </Link>
    </div>
  );
};
