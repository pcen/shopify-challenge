import { UserActionTypes } from '../actions/user';

// LOGGED_OUT_USER defines a logged out or invalid user
const LOGGED_OUT_USER = {
  loggedIn: false,
  username: '',
  authToken: '',
  id: null,
}

// getUserOnRefresh checks local storage, and returns the 'user'
// object if it exists and is logged in, otherwise it returns
// a logged out user object
const getUserOnRefresh = () => {
  let user = JSON.parse(localStorage.getItem('user'));
  return user && user.loggedIn ? user : LOGGED_OUT_USER;
}

// User Reducer
const userReducer = (state = getUserOnRefresh(), action) => {
  switch (action.type) {
    case UserActionTypes.LoginSuccess:
      // Set state.user to action.user upon successful login
      return action.user;
    case UserActionTypes.LoginFailure:
    case UserActionTypes.Logout:
      // Set state.user to LOGGED_OUT_USER upon logout or failed login
      return LOGGED_OUT_USER;
    default:
      return state;
  }
}

export default userReducer;
