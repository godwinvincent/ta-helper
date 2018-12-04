import React, { Component } from 'react'; //import React Component
import {FormGroup, Label, Input, Button } from 'reactstrap'
import {Redirect} from 'react-router-dom'

export default class EmailVerifyForm extends Component {
    constructor(props){
      super(props);
      this.state = {}
    }
  
    handleVerify(event) {
      event.preventDefault(); //don't submit
    //   this.props.signInCallback(this.state.email, this.state.password);
    }
    handleChange(event){
      let newState = {};
      newState[event.target.name] = event.target.value;
      this.setState(newState);
    }
  
    render() {
      return (this.props.redirect ? <Redirect to="/" /> :(
        <form>
          <FormGroup>
            <Label for="Email Verification Code">Email</Label>
            <Input onChange = {e => this.handleChange(e)} id="code" 
              type="code" 
              name="code"
              />
          </FormGroup>
          <FormGroup>
            <Button color="primary" onClick={(e) => this.handleVerify(e)} >
              Submit
            </Button>
          </FormGroup>
        </form>))
      
    }
  }