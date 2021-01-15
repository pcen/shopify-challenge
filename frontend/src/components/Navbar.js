import { Link, useLocation } from 'react-router-dom';

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

  if (noNavRoutes.has(location.pathname)) {
    return null;
  }

  return (
    <div className='navbar'>
      {Links.map((link, i) => {
        return link.path === location.pathname ?
          (<div key={i} className='current-link'>{link.name}</div>)
          :
          (<Link key={i} className='link' to={link.path}>{link.name}</Link>)
      })}
      <LogoutButton />
    </div>
  );
}

export default Navbar;
