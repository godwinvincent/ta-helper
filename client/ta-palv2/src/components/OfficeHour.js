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
        fetch("https://tapalapi.patrickold.me/v1/officehours/?oh=" + this.state.id, {
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
    fetch("https://tapalapi.patrickold.me/v1/officehours/?oh=" + this.state.id, {
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
        var question = Buffer.from(data, 'base64').toString('ascii')
        console.log(question)
        this.update();
      }

    deleteQuestion(id) {
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

    changeQuestionOrder(change, qID) {
        var bodyObj = {
            "mode" : "order",
            "update" : change
        }
        var auth = localStorage.getItem('Authorization');
        fetch("https://tapalapi.patrickold.me/v1/question/?qid="+qID, {
            method: "PATCH",
            mode: "cors", 
            headers: {
                "Authorization": auth,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(bodyObj)
        })
        .then(response => {
            if (response.status < 300) {
                // this.setState({question:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            // error.text().then(error => this.setState({ errorquestion: error }))
            console.log(error)
        })
    }

    changeQuestionUsers(qID, operation) {
        var auth = localStorage.getItem('Authorization');
        var methodToUse = ""
        if (operation === "add") {
            methodToUse = "POST"
        } else if (operation === "remove") {
            methodToUse = "DELETE"
        }
        fetch("https://tapalapi.patrickold.me/v1/question/?qid="+qID, {
            method: methodToUse,
            headers: {
                "Authorization": auth,
                "Content-Type": "application/json",
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

    sendNotification(students) {
        // notification code goes here
    }

    render() {
        var userPull = JSON.parse(localStorage.getItem("User"))
        return ( this.props.currentUser ?
        (
            <div>
                <Header showOptions={false}/>
                <QuestionList questions={this.state.questions} currentUser={this.props.currentUser} 
                deleteQuestionCallback={(id) => this.deleteQuestion(id)}  editQuestionCallback={(id, msg) => this.editQuestion(id, msg)} 
                changeQuestionOrder={(change, qID) => this.changeQuestionOrder(change, qID)} 
                changeQuestionUsers={(qID, operation) => this.changeQuestionUsers(qID, operation)} 
                sendNotification={(students) => this.sendNotification(students)} id={this.state.id} />
                { userPull.role == "student" ?
                <QuestionBox currentUser={this.props.currentUser} id={this.state.id} /> : ""}
                <Websocket url={'wss://tapalapi.patrickold.me/v1/ws?auth=' + localStorage.getItem('Authorization')}
              onquestion={this.handleData.bind(this)}/>
            </div>
        ) :
        <Redirect to="/"/>
        )
    }
}