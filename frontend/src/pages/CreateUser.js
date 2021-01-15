import React from 'react';
import { Link } from 'react-router-dom';

import CreateUserForm from '../components/CreateUserForm';

import '../styles/create_user.css';

// Login Page
const CreateUser = props => {
  return (
    <React.Fragment>
      <div className='create-user'>
        <h1>Create Account</h1>
        <CreateUserForm />
      </div>
      <br></br>
      <Link className='create-user-link' to='/'>Back to Login</Link>
    </React.Fragment>
  );
}

export default CreateUser;
