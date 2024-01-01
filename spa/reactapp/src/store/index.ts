import { configureStore } from '@reduxjs/toolkit';
import loginReducer from '../features/login/LoginSlice';

// RootStateという型を定義します。これはアプリの全てのステート型を含むものです。
export type RootState = {
  login: ReturnType<typeof loginReducer>;
};

// Storeを設定します。Reducerの型もTypeScriptが推論できるようになります。
export const store = configureStore({
  reducer: {
    login: loginReducer,
  },
});
