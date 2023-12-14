import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';

interface AuthState {
  isLoggedIn: boolean;
  user: {
    email: string;
  } | null;
  error: string | null;
}

const initialState: AuthState = {
  isLoggedIn: false,
  user: null,
  error: null,
};

export const loginUser = createAsyncThunk(
  'auth/loginUser',
  async ({ email, password }: { email: string; password: string }, { rejectWithValue }) => {
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const data = await response.json();
      return { email: data.email };
    } catch (error) {
      return rejectWithValue('Invalid email or password');
    }
  }
);

export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout: (state) => {
      state.isLoggedIn = false;
      state.user = null;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(loginUser.fulfilled, (state, action: PayloadAction<{ email: string }>) => {
      state.isLoggedIn = true;
      state.user = { email: action.payload.email };
      state.error = null;
    });
    builder.addCase(loginUser.rejected, (state, action) => {
      state.isLoggedIn = false;
      state.user = null;
      state.error = action.payload as string;
    });
  },
});

export const { logout } = authSlice.actions;

export default authSlice.reducer;
