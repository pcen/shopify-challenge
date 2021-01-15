import React from 'react';
import LoginForm from '../components/LoginForm';
import { Link } from 'react-router-dom';

import '../styles/login.css';

// Login Page
const Login = props => {
  return (
    <React.Fragment>
      <div className='login'>
        <h1>Login</h1>
        <LoginForm />
      </div>
      <br></br>
      <Link className='create-user-link' to='/create-user'>Create Account</Link>
    </React.Fragment>
  )
}

export default Login;
