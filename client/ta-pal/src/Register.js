import React, { Component } from 'react';
import './Register.css'
//import {Link} from "react-router-dom";

class Register extends Component {
    constructor(props) {
      super(props);
      this.state = {
        userRole: 'student',
      };
    }

    onSubmitClick = () => {

    }

    render() {
        // obtain different info depending on if student or instructor
        if (this.state.userRole === 'student') {
            return (
                <div className="login-block" id="block">
                    <div className="container">
                        <div className="row">
                            <div className="col-md-4 login-sec">
                                <h2 className="text-center">Register</h2>
                                <form className="login-form">
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">Email Address</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-email"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Username</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-username"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">First Name</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-firstname"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Last Name</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-lastname"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">Password</label>
                                        <input type="password" className="form-control" placeholder="" id="reg-password"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Confirm Password</label>
                                        <input type="password" className="form-control" placeholder="" id="reg-passwordconf"></input>
                                    </div>
                                    <button type="button" onClick={this.onSubmitClick()} className="btn btn-primary">Submit</button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            );
        } else if (this.state.userRole === 'instructor') {
            return (
                <div className="login-block" id="block">
                    <div className="container">
                        <div className="row">
                            <div className="col-md-4 login-sec">
                                <h2 className="text-center">Register</h2>
                                <form className="login-form">
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">Email Address</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-email"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Username</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-username"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">First Name</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-firstname"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Last Name</label>
                                        <input type="text" className="form-control" placeholder="" id="reg-lastname"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputEmail1" className="text-uppercase">Password</label>
                                        <input type="password" className="form-control" placeholder="" id="reg-password"></input>
                                    </div>
                                    <div className="form-group">
                                        <label htmlFor="exampleInputPassword1" className="text-uppercase">Confirm Password</label>
                                        <input type="password" className="form-control" placeholder="" id="reg-passwordconf"></input>
                                    </div>
                                    <button type="button" onClick={this.onSubmitClick()} className="btn btn-primary">Submit</button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            )
        }
    }
}

export default Register;