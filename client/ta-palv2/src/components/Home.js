import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Header from './Header';
import OfficeHourList from './OfficeHourList';
import { Card, CardImg, CardText, CardBody, CardTitle, CardSubtitle, Button, Container, Row, Col } from 'reactstrap';
import EmailVerifyForm from './EmailVerifyForm'

export default class Home extends Component {

    constructor(props) {
        super(props);
        this.state = {
            channels: {}
        };
    }

    componentDidMount(){
    }

    componentWillUnmount(){

    }

    postNewOfficeHours(name) {
        var officeHourData = {
            "name": name
        }
        var auth = localStorage.getItem('Authorization');
        fetch("http://localhost:80/v1/officehours", {
            method: "POST",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(officeHourData), 
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({message:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => console.log(error))
        })
    }

    editChannel(id, channel, desc) {
        var channelData = {
            "name": channel,
            "description": desc
        }
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/channels/"+id, {
            method: "PATCH",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(channelData), 
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({message:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorMessage: error }))
        })
    }

    editUserChannel(channelID, userID, add) {
        var channelData = {
            "userID": userID
        }
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/channels/"+ channelID + "/members", {
            method: add ? "POST" : "DELETE",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(channelData), 
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({message:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorMessage: error }))
        })
    }

    deleteChannel(channelID) {
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/channels/"+ channelID, {
            method: "DELETE",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            }
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({message:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorMessage: error }))
        })
    }

    render() {
        let content = "";
        if(this.props.loading){
             content = (<div className="text-center"><i className="fa fa-spinner fa-spin fa-3x" aria-label="Connecting..."></i></div>)
        }
        else{
            content = this.props.user ?
                this.props.user.emailActivated ?
                <div>
                    <Header newOfficeHourCallback={(name) => this.postNewOfficeHours(name)} signOutCallback={this.props.signOutCallback} showOptions={true} />
                    <OfficeHourList deleteChannelCallback={(channelID) => this.deleteChannel(channelID)} editChannelUserCallback={(channelID, userID, add) => this.editUserChannel(channelID, userID, add)} editChannelCallback={(id, channel, desc) => this.editChannel(id, channel, desc)} user={this.props.user} ref={this.ref} path="channelsList/" redirect="/channels/" />
                </div> :
                <EmailVerifyForm />
            :
                (<Container>
                    <Row>
                        <Col sm="12" md={{ size: 8, offset: 2 }}>
                            <Card>
                                <CardImg top width="300" height="400" src="img//tight.jpg" alt="Card image cap" />
                                <CardBody>
                                    <CardTitle className="text-center">Welcome to TA-Pal!</CardTitle>
                                    <CardSubtitle className="text-center">Making Office Hours Better!</CardSubtitle>
                                    <CardText className="text-center">Please Log In or Sign up to get Started</CardText>
                                    <div className="text-center">
                                        <Link to="/login"><Button>Log In</Button></Link>
                                        <Link to="/join"><Button>Sign Up</Button></Link>
                                    </div>
                                </CardBody>
                            </Card>
                        </Col>
                    </Row>
                </Container>)
        }
        return (content)
    }
}