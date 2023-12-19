import { configureStore, createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import axios from 'axios';

// ログインリクエストのための型定義
interface LoginRequest {
  email: string;
  password: string;
}

// レスポンスの型定義
interface LoginResponse {
  token: string;
}

// 初期状態の型定義
interface LoginState {
  email: string | null;
  isLoggedIn: boolean;
  error: string | null;
  token: string | null;
}

const initialLoginState: LoginState = {
  email: null,
  isLoggedIn: false,
  error: null,
  token: null,
};

// 非同期アクションの定義
export const loginUserAsync = createAsyncThunk(
  'login/loginUserAsync',
  async ({ email, password }: LoginRequest, { dispatch }) => {
    // 成功した場合、emailを保存
    dispatch(setEmail(email));
    const response = await axios.post<LoginResponse>('http://localhost:8080/login', {
      email,
      password,
    });
    return response.data;
  }
);

const loginSlice = createSlice({
  name: 'login',
  initialState: initialLoginState,
  reducers: {
    setEmail: (state, action: PayloadAction<string>) => {
      state.email = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(loginUserAsync.fulfilled, (state, action: PayloadAction<LoginResponse>) => {
      state.isLoggedIn = true;
      state.token = action.payload.token;
      state.error = null;
    });
  },
});

const rootReducer = {
  login: loginSlice.reducer,
};

const store = configureStore({
  reducer: rootReducer,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export const { setEmail } = loginSlice.actions;
export default store;
