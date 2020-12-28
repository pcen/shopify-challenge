import './App.css';
import React, { useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { useStore } from 'react-redux';

import Routes, { Login, AuthorizedRoute } from './Routes';

const getBackendStatus = () => {
  let endpoint = '/status';
  return fetch(endpoint).then(r => r.text());
}

const App = props => {
  const store = useStore();

  const handleClick = () => {
    let user = store.getState().user;
    fetch('/images', {
      method: 'GET',
      headers: { 'Authorization': user.authToken, },
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
      <BrowserRouter>
        <Switch>
          <Route exact path='/' component={Login} />
          {Routes.map(r => {
            return (
              <AuthorizedRoute
                key={r.path}
                path={r.path}
                component={r.component}
                useAuth={r.useAuth}
              />
            )
          })}
        </Switch>
      </BrowserRouter>

      <br />
      <br />
      <button onClick={handleClick}>
        Get Image Data
      </button>
    </div>
  );
}

export default App;
