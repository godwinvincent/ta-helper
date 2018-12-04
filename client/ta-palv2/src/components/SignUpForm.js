import React, { Component } from 'react'; //import React Component
import {FormGroup, Label, Input, Button} from 'reactstrap'
import {Redirect } from 'react-router-dom'

export default class SignUpForm extends Component {
    constructor(props){
      super(props);
      this.state = {
        email : undefined,
        password : undefined,
        passwordConf : undefined,
        username: undefined,
        firstName: undefined,
        lastName: undefined
      }; //initialize state
    }
  
    //handle signUp button
    handleSignUp(event) {
      event.preventDefault(); //don't submit
      this.props.signUpCallback(this.state.email, this.state.password, this.state.passwordConf, this.state.username, this.state.firstName, this.state.lastName);
    }
  
    handleChange(event){
      let newState = {};
      newState[event.target.name] = event.target.value;
      this.setState(newState);
    }
  
    /* SignUpForm#render() */
    render() {
        return (this.props.redirect ? <Redirect to="/" /> :(
        <form>
          {/* email */}
          <FormGroup>
            <Label for="email">Email</Label>
            <Input onChange = {e => this.handleChange(e)} id="email" 
              type="email" 
              name="email"
              />
          </FormGroup>
          
          {/* password */}
          <FormGroup>
            <Label for="password">Password</Label>
            <Input onChange = {e => this.handleChange(e)} id="password" 
              type="password"
              name="password"
              />
          </FormGroup>

          {/* password confirmation */}
          <FormGroup>
            <Label for="passwordConf">Password Confirmation</Label>
            <Input onChange = {e => this.handleChange(e)} id="passwordConf" 
              type="password"
              name="passwordConf"
              />
          </FormGroup>
  
          {/* username */}
          <FormGroup>
            <Label htmlFor="username">Username</Label>
            <Input onChange = {e => this.handleChange(e)} id="username" 
              name="username"
              />
          </FormGroup>
  
          {/* firstname */}
          <FormGroup>
            <Label htmlFor="firstname">First Name</Label>
            <Input onChange = {e => this.handleChange(e)} id="firstName" 
              name="firstName"
              />
          </FormGroup>

           {/* lastName */}
           <FormGroup>
            <Label htmlFor="lastName">Username</Label>
            <Input onChange = {e => this.handleChange(e)} id="lastName" 
              name="lastName"
              />
          </FormGroup>
  
          {/* buttons */}
          <FormGroup>
            <Button color="primary" className="mr-2" onClick={(e) => this.handleSignUp(e)} >
              Sign-up
            </Button>
          </FormGroup>
        </form>))
      
    }
  }