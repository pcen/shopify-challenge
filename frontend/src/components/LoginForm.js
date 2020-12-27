import React, { useState } from 'react';
import { useDispatch } from 'react-redux';

import { login } from '../state/actions/user';

// LoginForm accepts a username and password as input, and will
// send the entered credentials to the backend upon form submission
const LoginForm = props => {
  const dispatch = useDispatch();

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [prev, setPrev] = useState({ username: '', password: '' });

  // True if the current form data is not the most recently submitted data
  const dataIsNew = () => username !== prev.username || password !== prev.password;
  // True if both username and password are not empty
  const dataIsValid = () => username.length !== 0 && password.length !== 0;

  const submitData = () => {
    if (dataIsNew()) {
      setPrev({ username, password });
      if (dataIsValid()) {
        dispatch(login(username, password));
      }
    }
  }

  const handleKeyPress = key => {
    if (key.code === 'Enter') { submitData(); }
  }

  return (
    <React.Fragment>
      <h1>Login</h1>
      <form onSubmit={submitData} onKeyPress={handleKeyPress} >
        <label htmlFor='uname'>Username:</label>
        <br />
        <input type='text' name='uname' onChange={e => setUsername(e.target.value)} />
        <br />
        <label htmlFor='pword'>Password:</label><br />
        <input type='password' name='pname' onChange={e => setPassword(e.target.value)} />
      </form>
    </React.Fragment>
  )
}

export default LoginForm;
