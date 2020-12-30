import './App.css';
import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';

import Routes, { Login, AuthorizedRoute } from './Routes';

const App = props => {

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
    </div>
  );
}

export default App;
