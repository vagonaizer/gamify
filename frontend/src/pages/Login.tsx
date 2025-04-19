import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import apiClient from '../api/client';

const Login: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const { data } = await apiClient.post('/login', { email, password });
      localStorage.setItem('user', JSON.stringify(data));
      navigate('/home');
    } catch {
      setError('Invalid credentials');
    }
  };

  return (
    <div className="authContainer">
      <h1>Login</h1>
      <form className="form" onSubmit={handleSubmit}>
        <div className="formGroup">
          <label htmlFor="login-email">Email</label>
          <input
            id="login-email"
            type="email"
            placeholder="Enter your email"
            title="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="login-password">Password</label>
          <input
            id="login-password"
            type="password"
            placeholder="Enter your password"
            title="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Log In</button>
        {error && <p className="error">{error}</p>}
      </form>
      <p className="signup-link">
        Don’t have an account? <a href="/register">Sign Up</a>
      </p>
    </div>
  );
};

export default Login;