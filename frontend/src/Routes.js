import { Route, Redirect } from 'react-router-dom';
import { useStore } from 'react-redux';

// Import page components
import Home from './pages/Home';
import Login from './pages/Login';


const AuthorizedRoute = props => {
  const { path, component, useAuth } = props;
  const user = useStore().getState().user;
  if (useAuth && user.loggedIn === false) {
    return <Redirect to='/' />
  } else {
    return <Route path={path} component={component} />
  }
}

const Routes = [
  {
    path: '/home',
    component: Home,
    useAuth: true,
  },
]

export { AuthorizedRoute, Login };
export default Routes;