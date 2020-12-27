// User Actions

import { postJSON } from '../../utils/requests';

const login = (username, password) => {

  const success = user => {
    return { type: 'LOGIN_SUCCESS', user, };
  }

  const failure = error => {
    return { type: 'LOGIN_FAILURE', error, };
  }

  return dispatch => {
    postJSON('/login', { username, password }).then(r => {
      if (r.success) {
        dispatch(success(r.user));
      } else {
        dispatch(failure(r.error));
      }
    })
  }
}

export { login };
