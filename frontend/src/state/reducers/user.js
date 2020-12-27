// User Reducer

const INITIAL_USER_STATE = {
  Username: '',
  Password: '',
  LoggedIn: false,
  JWT: '',
};

const userReducer = (state = INITIAL_USER_STATE, action) => {
  switch (action.type) {
    case 'LOGIN_SUCCESS':
      return { loggedIn: true };
    case 'LOGIN_FAILURE':
      return { loggedIn: false };
    default:
      return state;
  }
}

export default userReducer;
