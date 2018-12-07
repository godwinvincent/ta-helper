import React, { Component } from 'react'; //import React Component
import {FormGroup, Label, Input, Button, Row, Col } from 'reactstrap'
import {Redirect} from 'react-router-dom'
import './styles/Fields.css'

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

    handleGoBack(event) {
      event.preventDefault();
      // window.location = '/'
      this.setState({redirect : true})
    }

    // go back to home page
    handleSignOut(event) {
      event.preventDefault();
      this.props.signOutCallback();
    }

    handleChange(event){
      let newState = {};
      newState[event.target.name] = event.target.value;
      this.setState(newState);
    }
  
    /* SignUpForm#render() */
    render() {
      return (this.props.redirect || this.state.redirect  ? <Redirect to="/" /> :(
        <form style={{marginTop: '20px'}}>
          <Row id="email-row">
            <Col sm={{ size: 6, offset: 3 }}>
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
                <Button color="warning" onClick={(e) => this.handleGoBack(e)} style={{float: 'right'}} >
                  Back
                </Button>
              </FormGroup>
            </Col>
          </Row>
        </form>))
      
    }
  }