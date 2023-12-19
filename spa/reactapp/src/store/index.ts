import { configureStore, createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import axios, { AxiosError } from 'axios';

// ログインリクエストのための型定義
interface LoginRequest {
  email: string;
  password: string;
}

// レスポンスの型定義
interface LoginResponse {
  email: string;
  token: string;
}

// エラーレスポンスの型定義
interface ErrorResponse {
  message: string;
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
export const loginUserAsync = createAsyncThunk<LoginResponse, LoginRequest, { rejectValue: ErrorResponse }>(
  'login/loginUserAsync',
  async ({ email, password }, { rejectWithValue }) => {
    try {
      const response = await axios.post<LoginResponse>('http://localhost:8080/login', {
        email,
        password,
      });
      return response.data;
    } catch (error) {
      let errResponse = (error as AxiosError<ErrorResponse>).response;
      return rejectWithValue(errResponse?.data || { message: 'Unknown error occurred' });
    }
  }
);

// スライスの定義
const loginSlice = createSlice({
  name: 'login',
  initialState: initialLoginState,
  reducers: {
    // ここに必要な同期的なリデューサーを追加することができます
  },
  extraReducers: (builder) => {
    builder.addCase(loginUserAsync.fulfilled, (state, action: PayloadAction<LoginResponse>) => {
      // ログインが成功した場合の状態更新
      state.email = action.payload.email;
      state.isLoggedIn = true;
      state.token = action.payload.token;
      state.error = null;
    });
    builder.addCase(loginUserAsync.rejected, (state, action: PayloadAction<ErrorResponse | undefined>) => {
      // ログインが失敗した場合の状態更新
      state.error = action.payload?.message || 'Login failed';
    });
  },
});

// ストアの設定
const store = configureStore({
  reducer: {
    login: loginSlice.reducer,
  },
});

// 型定義のエクスポート
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
