import { Route, Redirect } from 'react-router-dom';
import { useStore } from 'react-redux';

// Import page components
import CreateUser from './pages/CreateUser';
import Home from './pages/Home';
import Login from './pages/Login';
import Upload from './pages/Upload';

const AuthorizedRoute = props => {
  const { path, component, useAuth } = props;
  const user = useStore().getState().user;
  if (useAuth === true && user.loggedIn === false) {
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
  {
    path: '/upload',
    component: Upload,
    useAuth: true,
  },
  {
    path: '/create-user',
    component: CreateUser,
    useAuth: false,
  }
]

export { AuthorizedRoute, Login };
export default Routes;
