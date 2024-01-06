import { Link } from 'react-router-dom';

export const UserDetail = () => {
  // CSSを記述する
  const textStyle = {
    color: 'blue',
    fontWeight: 'bold',
  };

  return (
    <div>
      <h1>User Details</h1>
      <p>User ID: </p>
      <Link to="/" style={textStyle}>
        Homeに戻る
      </Link>
    </div>
  );
};
