import { configureStore } from '@reduxjs/toolkit';
import authReducer from './authSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
  },
});

// この型は、useSelectorでの使用のために必要です。
export type RootState = ReturnType<typeof store.getState>;

// この型は、useDispatchのために必要です。
export type AppDispatch = typeof store.dispatch;
