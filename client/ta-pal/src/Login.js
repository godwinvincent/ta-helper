import React, { Component } from 'react';
import './Login.css'
//import {Link} from "react-router-dom";

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
                <h1>
                    Welcome to the Office Hour Manager!
                </h1>
                <div>
                    <button className="btn btn-primary btn-lg" id="student-button" onClick={this.setRedirect}>
                        Student View
                    </button>
                </div>
                <div>
                    <button className="btn btn-primary btn-lg" id="ta-button">
                        Teaching Assistant View
                    </button>
                </div>
                <footer>
                    Created by Ben, Godwin, and Patrick
                </footer>
            </div>
        );
    }
}

export default Login;