import React, { Component } from 'react';
import './Login.css'
class Login extends Component {
    constructor(props) {
      super(props);
      this.state = {
        value: null,
      };
    }
    render() {
        return (
            <div className="Login">
                <h2>
                    Welcome to the Office Hour Manager!
                </h2>
                <button id="student-button">
                    Student View
                </button>
                <button id="ta-button">
                    Teaching Assistant View
                </button>
                <footer>
                    Created by Ben, Godwin, and Patrick
                </footer>
            </div>
        );
    }
}

export default Login;