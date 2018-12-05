import React, { Component } from 'react'; //import React Component
import {FormGroup, Label, Input, Button } from 'reactstrap'
import {Redirect} from 'react-router-dom'

export default class SignInForm extends Component {
    constructor(props){
      super(props);
      this.state = {
        email : undefined,
        password : undefined
      };
    }
  
    handleSignIn(event) {
      event.preventDefault(); //don't submit
      this.props.signInCallback(this.state.netid, this.state.password);
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
          <FormGroup>
            <Label for="netid">Net ID</Label>
            <Input onChange = {e => this.handleChange(e)} id="netid" 
              type="netid" 
              name="netid"
              />
          </FormGroup>
          <FormGroup>
            <Label for="password">Password</Label>
            <Input onChange = {e => this.handleChange(e)} id="password" 
              type="password"
              name="password"
              />
          </FormGroup>

          <FormGroup>
            <Button color="primary" onClick={(e) => this.handleSignIn(e)} >
              Sign-in
            </Button>
          </FormGroup>
        </form>))
      
    }
  }