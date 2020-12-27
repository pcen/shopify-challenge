import './App.css';
import React, { useEffect } from 'react';

import LoginForm from './components/LoginForm';

const getBackendStatus = () => {
  let endpoint = '/status';
  return fetch(endpoint).then(r => r.text());
}

const App = props => {

  useEffect(() => {
    getBackendStatus().then(text => {
      console.log(text);
    });
  }, []);

  return (
    <div className="App">
      <LoginForm />
    </div>
  );
}

export default App;
