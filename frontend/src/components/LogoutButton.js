import React from 'react';
import { Redirect } from 'react-router-dom';
import { useDispatch, connect } from 'react-redux';

import { logout } from '../state/actions/user';

import '../styles/navbar.css';

// LogoutButton logs the user out, removing their credentials from
// Redux store and browser storage
const LogoutButton = props => {
  const dispatch = useDispatch();

  const { loggedIn } = props;

  const handleLogout = () => {
    dispatch(logout());
  }

  // Redirect to login if the user is not logged in
  if (!loggedIn) {
    return (<Redirect to='/' />);
  } else {
    return (
      <div className='link' onClick={handleLogout}>
        Logout
      </div>
    )
  }
}

// Map user state to props so the logout button can redirect to
// the login page once the user is logged out
const mapState = state => {
  return { loggedIn: state.user.loggedIn };
}

// Create the connected LoginForm component
const ConnectedLogoutButton = connect(mapState)(LogoutButton);

export default ConnectedLogoutButton;
