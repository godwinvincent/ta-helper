import React, { Component } from 'react'; //import React Component
import {FormGroup, Label, Input, Button, Row, Col } from 'reactstrap'
import {Redirect} from 'react-router-dom'

export default class EmailVerifyForm extends Component {
    constructor(props){
      super(props);
      this.state = {}
    }
    handleSend(event){
      event.preventDefault();
      var user = localStorage.getItem("User")
      var auth = localStorage.getItem("Authorization")
      console.log(auth, user)
      fetch("https://tapalapi.patrickold.me/v1/email", {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                "Authorization": auth
            }
        })
        .then(response => {
            if (response.status < 300) {
              alert("Email sent succesfully!")
            } else {
                throw response
            }
        })
        .catch(function(error) {
          console.log(error)
            // error.text().then(error => alert("error"))
        })
    }
    handleVerify(event){
      event.preventDefault();
      var user = localStorage.getItem("User")
      var auth = localStorage.getItem("Authorization")
      console.log(auth, user)
      fetch("https://tapalapi.patrickold.me/v1/email/verify?c=" + this.state.code, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                "Authorization": auth
            }
        })
        .then(response => {
            if (response.status < 300) {
                alert("email verified!, please log in again!")
                localStorage.removeItem("Authorization")
                localStorage.removeItem("User")
                window.location.reload(); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            //error.text().then(error => alert("Error verifying"))
        })
    }

    handleGoHome(event) {
      event.preventDefault();
      localStorage.removeItem("Authorization");
      localStorage.removeItem("User");
      window.location.reload(); 
      window.location = "/";
    }

    handleChange(event){
      let newState = {};
      newState[event.target.name] = event.target.value;
      this.setState(newState);
    }
  
    render() {
      return (this.props.redirect ? <Redirect to="/" /> :(
        <span>
        <form>
        <Row id="email-row">
            <Col sm={{ size: 6, offset: 3 }}>
              <FormGroup>
                <Label for="Email Verification Code">Email Verification Code</Label>
                <Input onChange = {e => this.handleChange(e)} id="code" 
                  type="code" 
                  name="code"
                  />
              </FormGroup>
              <FormGroup>
                <Button color="primary" onClick={(e) => this.handleSend(e)} >
                  Send Email
                </Button>
                <Button color="warning" style={{float: 'right'}} onClick={(e) => this.handleGoHome(e)} >
                  Home
                </Button>
              </FormGroup>
              <FormGroup>
                <Button color="primary" onClick={(e) => this.handleVerify(e)} >
                  Submit
                </Button>
              </FormGroup>
            </Col>
          </Row>
        </form> 
        </span>))
    }
  }