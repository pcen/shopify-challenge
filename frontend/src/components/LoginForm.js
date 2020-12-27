import React, { useState } from 'react';
import { postJSON } from '../utils/requests';

// LoginForm accepts a username and password as input, and will
// send the entered credentials to the backend upon form submission
const LoginForm = props => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [prev, setPrev] = useState({ username: '', password: '' });

  const dataIsNew = () => {
    return username !== prev.username || password !== prev.password;
  }

  const submitData = event => {
    if (dataIsNew()) {
      setPrev({ username, password });
      console.log(`Sending login: Username: ${username}\nPassword: ${password}`);
      postJSON('/login', { Username: username, Password: password }).then(json => {
        console.log(JSON.stringify(json));
      })
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
