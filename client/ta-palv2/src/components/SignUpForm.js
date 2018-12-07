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
      console.log(this.state.firstName)
      event.preventDefault(); //don't submit
      this.props.signUpCallback(this.state.netid, this.state.password, this.state.passwordConf, this.state.firstName, this.state.lastName);
    }
  
    handleChange(event){
      let newState = {};
      newState[event.target.name] = event.target.value;
      this.setState(newState);
    }

    handleGoBack(event) {
      event.preventDefault();
      window.location = "/"
    }
  
    /* SignUpForm#render() */
    render() {
        return (this.props.redirect ? <Redirect to="/" /> : (
        <form style={{marginTop: '20px'}}>
            <Button color="warning" style={{marginBottom: '10px'}} onClick={(e) => this.handleGoBack(e)} >
              Back
            </Button>
          {/* email */}
          <FormGroup>
            <Label for="email">UW NET ID</Label>
            <Input onChange = {e => this.handleChange(e)} id="netid" 
              type="netid" 
              name="netid"
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
  
          {/* firstname */}
          <FormGroup>
            <Label htmlFor="firstname">First Name</Label>
            <Input onChange = {e => this.handleChange(e)} id="firstName" 
              name="firstName"
              />
          </FormGroup>

           {/* lastName */}
           <FormGroup>
            <Label htmlFor="lastName">Last Name</Label>
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