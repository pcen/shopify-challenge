import './App.css';
import React, { useEffect } from 'react';
import { useStore } from 'react-redux';

import LoginForm from './components/LoginForm';

const getBackendStatus = () => {
  let endpoint = '/status';
  return fetch(endpoint).then(r => r.text());
}

const App = props => {
  const store = useStore();

  const handleClick = () => {
    let user = store.getState().user;
    // console.log('current user:', user);
    fetch('/images', {
      method: 'GET',
      headers: {
        'Authorization': user.authToken,
      },
    }).then(r => r.json()).then(j => {
      console.log(j);
    })
  }

  useEffect(() => {
    getBackendStatus().then(text => {
      console.log(text);
    });
  }, []);

  return (
    <div className="App">
      <LoginForm />
      <br /><br />
      <button onClick={handleClick}>
        Get Image Data
      </button>
    </div>
  );
}

export default App;
