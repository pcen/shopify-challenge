// User Actions

import { postJSON } from '../../utils/requests';

const UserActionTypes = {
  LoginSuccess: 'LOGIN_SUCCESS',
  LoginFailure: 'LOGIN_FAILURE',
  Logout: 'LOGOUT',
}

const login = (username, password) => {

  const success = user => {
    return { type: UserActionTypes.LoginSuccess, user, };
  }

  const failure = error => {
    return { type: UserActionTypes.LoginFailure, error, };
  }

  return dispatch => {
    postJSON('/login', { username, password }).then(r => {
      if (r.success) {
        // set 'loggedIn' to true
        r.user.loggedIn = true;
        localStorage.setItem('user', JSON.stringify(r.user));
        dispatch(success(r.user));
      } else {
        localStorage.removeItem('user');
        dispatch(failure(r.error));
      }
    })
  }
}

const logout = () => {
  localStorage.removeItem('user');
  return { type: UserActionTypes.Logout };
}

export { UserActionTypes, login, logout };
