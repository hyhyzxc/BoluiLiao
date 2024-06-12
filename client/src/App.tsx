import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Route, Routes } from 'react-router-dom';
import LoginPage from './components/LoginPage';

function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<LoginPage/>}></Route>
      </Routes>
    </div>
  );
}

export default App;
