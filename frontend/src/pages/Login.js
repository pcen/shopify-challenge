import LoginForm from '../components/LoginForm';

import '../styles/login.css';

// Login Page
const Login = props => {
  return (
    <div className='login'>
      <h1>Login</h1>
      <LoginForm />
    </div>
  )
}

export default Login;
