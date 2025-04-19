import React from 'react';

const HomePage: React.FC = () => {
  const user = JSON.parse(localStorage.getItem('user') || '{}');

  return (
    <div className="homeContainer">
      <h1>Welcome, {user.nickname || 'User'}</h1>
      <p>You are logged in!</p>
    </div>
  );
};

export default HomePage;