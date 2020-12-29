import './App.css';
import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';
import { useStore } from 'react-redux';

import Routes, { Login, AuthorizedRoute } from './Routes';
import { get } from './utils/requests';

const App = props => {
  const store = useStore();

  const handleClick = () => {
    let user = store.getState().user;
    get('/images', user.authToken).then(json => {
      console.log(json);
    })
  }

  return (
    <div className="App">
      <BrowserRouter>
        <Switch>
          {/* Login page */}
          <Route exact path='/' component={Login} />
          {/* Map route definitions in Routes.js */}
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
          {/* Undefined routes redirect to home page */}
          <Route>
            <Redirect to='/home' />
          </Route>
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
