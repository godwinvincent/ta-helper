import React, { Component } from 'react';
import QuestionBox from './QuestionBox'
import QuestionList from './QuestionList'
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
              this.setState({ questions: response });
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
            this.setState({ questions: response });
        })
        .catch(function(error) {
            error.text().then(error => alert("error"))
        })
    }

    editQuestion(id, msg) {
        var msgData = {
           "body": msg
        }
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/questions/"+id, {
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
                // this.setState({question:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorquestion: error }))
        })
    }

    handleData(data) {
        // var question = JSON.parse(Buffer.from(data, 'base64').toString('ascii'))
        this.update();
      }

    deletequestion(id) {
        var auth = localStorage.getItem('Authorization');
        fetch("https://info441api.godwinv.com/v1/questions/"+id, {
            method: "DELETE",
            mode: "cors", 
            headers: {
                "Authorization": auth
            }
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({question:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorquestion: error }))
        })
    }

    render() {
        return ( this.props.currentUser ?
        (
            <div>
                <Header showOptions={false}/>
                <questionList questions={this.state.questions} currentUser={this.props.currentUser} deletequestionCallback={(id) => this.deletequestion(id)}  editquestionCallback={(id, msg) => this.editquestion(id, msg)} id={this.state.id} />
                <questionBox currentUser={this.props.currentUser} id={this.state.id} />
                <Websocket url={'wss://info441api.godwinv.com/v1/ws?auth=' + localStorage.getItem('Authorization')}
              onquestion={this.handleData.bind(this)}/>
            </div>
        ) :
        <Redirect to="/"/>
        )
    }

}