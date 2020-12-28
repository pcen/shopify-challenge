import './App.css';
import React, { useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { useStore } from 'react-redux';

import Routes, { Login } from './Routes';

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
      <Router>
        <Switch>
          <Route exact path='/' component={Login} />
          {
            Routes.map(r => <Route path={r.path} component={r.page} />)
          }
        </Switch>
      </Router>

      <br /><br />
      <button onClick={handleClick}>
        Get Image Data
      </button>
    </div>
  );
}

export default App;
