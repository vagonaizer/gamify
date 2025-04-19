import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import apiClient from '../api/client';

const Register: React.FC = () => {
  const [email, setEmail] = useState('');
  const [nickname, setNickname] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const { data } = await apiClient.post('/register', { email, nickname, password });
      localStorage.setItem('user', JSON.stringify(data));
      navigate('/home');
    } catch {
      setError('Registration failed');
    }
  };

  return (
    <div className="authContainer">
      <h1>Register</h1>
      <form className="form" onSubmit={handleSubmit}>
        <div className="formGroup">
          <label htmlFor="register-email">Email</label>
          <input
            id="register-email"
            type="email"
            placeholder="Enter your email"
            title="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="register-nickname">Nickname</label>
          <input
            id="register-nickname"
            type="text"
            placeholder="Choose a nickname"
            title="Nickname"
            value={nickname}
            onChange={(e) => setNickname(e.target.value)}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="register-password">Password</label>
          <input
            id="register-password"
            type="password"
            placeholder="Enter your password"
            title="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Sign Up</button>
        {error && <p className="error">{error}</p>}
      </form>
      <p className="signup-link">
        Already have an account? <a href="/login">Log In</a>
      </p>
    </div>
  );
};

export default Register;