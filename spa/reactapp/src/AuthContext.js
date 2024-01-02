import React, { createContext, useState, useContext } from 'react';

const AuthContext = createContext(null);

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
  const [authToken, setAuthToken] = useState('');

  const updateToken = (token) => {
    setAuthToken(token);
  };

  return <AuthContext.Provider value={{ authToken, updateToken }}>{children}</AuthContext.Provider>;
};
