import { createSlice } from '@reduxjs/toolkit';

// userLoginにはAPIからのレスポンスが入っている
// パスワードはログイン時のみ使用するため、stateには含めない
const initialState = {
  email: '',
  authToken: '',
};

// payloadとして、emailとauthTokenを受け取る
const loginSlice = createSlice({
  name: 'user-login',
  initialState,
  reducers: {
    checkLoginUser: (state, action) => {
      // action.payloadからemailとauthTokenを抽出してstateに保存
      state.email = action.payload.email;
      state.authToken = action.payload.authToken;
    },
  },
});

// ここでdispatchのためにexportをしている。
export const { checkLoginUser } = loginSlice.actions;
export default loginSlice.reducer;
