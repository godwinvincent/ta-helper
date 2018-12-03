import React, { Component } from 'react';
import './App.css';
import Register from './Register.js'
import Login from './Login.js';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Register />
      </div>
    );
  }
}

export default App;
