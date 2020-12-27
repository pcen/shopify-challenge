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
      let action = r.success ? success : failure;
      dispatch(action(r.message));
    })
  }
}

export { login };
