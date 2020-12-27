// User Reducer

const USER_LOGGED_OUT = {
  loggedIn: false,
  username: '',
  authToken: '',
};

const userReducer = (state = USER_LOGGED_OUT, action) => {
  switch (action.type) {
    case 'LOGIN_SUCCESS':
      return {
        loggedIn: true,
        username: action.user.username,
        authToken: action.user.authToken,
      };
    case 'LOGIN_FAILURE':
      return USER_LOGGED_OUT;
    default:
      return state;
  }
}

export default userReducer;
