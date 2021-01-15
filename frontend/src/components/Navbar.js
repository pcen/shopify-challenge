import { Link, useLocation } from 'react-router-dom';
import { useStore } from 'react-redux';

import LogoutButton from './LogoutButton';

import '../styles/navbar.css';

const Links = [
  {
    name: 'Home',
    path: '/home',
  },
  {
    name: 'Upload',
    path: '/upload',
  },
]

// Set of routes that do not have a navbar
const noNavRoutes = new Set(['/', '/create-user']);

// Navbar component provides links to each page site.
const Navbar = props => {
  const location = useLocation();
  const user = useStore().getState().user;

  if (noNavRoutes.has(location.pathname)) {
    return null;
  }

  return (
    <div className='navbar'>
      <div className='navbar-username'>
        {user.username === null ? null :
          `Logged in as ${user.username}`
        }
      </div>
      <div className='navbar-link-container'>
        {Links.map((link, i) => {
          return link.path === location.pathname ?
            (<div key={i} className='current-link'>{link.name}</div>)
            :
            (<Link key={i} className='link' to={link.path}>{link.name}</Link>)
        })}
        <LogoutButton />
      </div>
    </div>
  );
}

export default Navbar;
