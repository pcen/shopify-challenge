import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { useDispatch, connect } from 'react-redux';

import { login } from '../state/actions/user';
import { postJSON } from '../utils/requests';

import '../styles/create_user.css';

// CreateUserForm accepts a username and password as input, and will
// send the entered credentials to the backend upon form submission
const CreateUserForm = props => {
  const dispatch = useDispatch();

  // loggedIn prop is mapped to Redux store for user data
  const { loggedIn } = props;

  // state variables for username, password, and most recently submitted data
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [prev, setPrev] = useState({ username: '', password: '' });

  // True if the current form data is not the most recently submitted data
  const dataIsNew = () => username !== prev.username || password !== prev.password;

  // True if both username and password are not empty
  const dataIsValid = () => username.length !== 0 && password.length !== 0;

  // Sends post request containing username and password to the backend
  const submitData = () => {

    if (dataIsNew()) {
      setPrev({ username, password });

      postJSON('/create-user', { username, password }).then(
        json => {
          dispatch(login(username, password));
        },
        error => {
          console.log(error);
        }
      );
    }
  }

  const handleKeyPress = key => {
    if (key.code === 'Enter') { submitData(); }
  }

  // If the user is logged in redirect to the home page
  if (loggedIn) {
    return (<Redirect to='/home' />);
  } else {
    return (
      <form className='create-user-form' onSubmit={submitData} onKeyPress={handleKeyPress} >
        <div>
          <label>Username:</label>
          <input type='text' name='uname' onChange={e => setUsername(e.target.value)} />
        </div>
        <div>
          <label>Password:</label>
          <input type='password' onChange={e => setPassword(e.target.value)} />
        </div>
        <div className='create-user-button' onClick={submitData}>Create Account</div>
      </form>
    )
  }
}

// Map user state to props so the login form can redirect to
// the home page once the user is authenticated
const mapState = state => {
  return { loggedIn: state.user.loggedIn };
}

// Create the connected CreateUserForm component
const ConnectedCreateUserForm = connect(mapState)(CreateUserForm);

export default ConnectedCreateUserForm;
