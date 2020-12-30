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

const Navbar = props => {
  const location = useLocation();

  if (location.pathname === '/') {
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
