import React, { Component } from 'react';
import MessageBox from './MessageBox'
import MessageList from './MessageList'
import Header from './Header'
import { Redirect } from 'react-router-dom'
import Websocket from 'react-websocket';


export default class OfficeHour extends Component {
    constructor(props) {
        super(props);
        this.state = {
            id: this.props.match.params.id
        };
    }
    componentDidMount(){
        var auth = localStorage.getItem('Authorization');
        fetch("http://localhost:80/v1/officehours/?oh=" + this.state.id, {
                method: "GET", // *GET, POST, PUT, DELETE, etc.
                mode: "cors", // no-cors, cors, *same-origin
                headers: {
                    "Authorization": auth
                }
            })
            .then(response => {
                if (response.status < 300) {
                    return response.json()
                } else {
                    throw response
                }
            })
            .then(response => {
              this.setState({ messages: response });
            })
            .catch(function(error) {
                error.text().then(error => alert("error"))
            })
      }

    update(){
    var auth = localStorage.getItem('Authorization');
    fetch("http://localhost:80/v1/officehours/?oh=" + this.state.id, {
            method: "GET", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                "Authorization": auth
            }
        })
        .then(response => {
            if (response.status < 300) {
                return response.json()
            } else {
                throw response
            }
        })
        .then(response => {
            this.setState({ messages: response });
        })
        .catch(function(error) {
            error.text().then(error => alert("error"))
        })
    }

    editMessage(id, msg) {
        var msgData = {
           "body": msg
        }
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/messages/"+id, {
            method: "PATCH",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(msgData), 
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

    handleData(data) {
        // var message = JSON.parse(Buffer.from(data, 'base64').toString('ascii'))
        this.update();
      }

    deleteMessage(id) {
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/messages/"+id, {
            method: "DELETE",
            mode: "cors", 
            headers: {
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
        return ( this.props.currentUser ?
        (
            <div>
                <Header showOptions={false}/>
                <MessageList messages={this.state.messages} currentUser={this.props.currentUser} deleteMessageCallback={(id) => this.deleteMessage(id)}  editMessageCallback={(id, msg) => this.editMessage(id, msg)} id={this.state.id} />
                <MessageBox currentUser={this.props.currentUser} id={this.state.id} />
                <Websocket url={'wss://info441api.godwinv.com/v1/ws?auth=' + localStorage.getItem('Authorization')}
              onMessage={this.handleData.bind(this)}/>
            </div>
        ) :
        <Redirect to="/"/>
        )
    }

}