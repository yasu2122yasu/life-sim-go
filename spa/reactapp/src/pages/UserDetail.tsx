import { useParams } from 'react-router-dom';

export const UserDetail = () => {
  const { userId } = useParams();

  // Fetch user data based on the id parameter

  return (
    <div>
      <h1>User Details</h1>

      <p>User ID: {userId}</p>
      {/* Display user information here */}
    </div>
  );
};
