import React, { Component } from 'react'; //import React Component
import './styles/Questions.css';

//A form the user can use to post a Chirp
export default class QuestionBox extends Component {
  constructor(props){
    super(props);
    this.state = {question:''};
  }

  post(body, type) {
    var questionData = {
        "questBody": body,
        "questType": type
    }
    var auth = localStorage.getItem('Authorization');
    console.log('this.props.id', this.props.id)
    fetch("https://tapalapi.patrickold.me/v1/officehours/?oh=" + this.props.id, {
            method: "POST",
            mode: "cors",
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(questionData), 
        })
        .then(response => {
            if (response.status < 300) {
                this.setState({question:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => 
              console.log(error)
              //this.setState( { errorMessage: error }))
        )})
  }

  updateQuestion(event) {
    console.log("updating")
    this.setState({question: event.target.value});
  }
  postQuestion(event, type){
    event.preventDefault(); 
    this.post(this.state.question, type)
    this.setState({question:''}); 
  }


  render() {
    return (
      <div className="container">
        <div className="row py-3 chirp-box">
          <div className="col pl-4 pl-lg-1">
            <form>
              <textarea name="text" className="form-control mb-2" placeholder="Type Question Here" 
                value={this.state.question} 
                onChange={(e) => this.updateQuestion(e)}
                />
              
              <div className="text-right">
                <button className="btn btn-primary" 
                  onClick={(e) => this.postQuestion(e, "short")} 
                  >
                  <i className="fa fa-pencil-square-o" aria-hidden="true"></i> Post Short Question
                </button>
                <button className="btn btn-primary m-2" 
                  onClick={(e) => this.postQuestion(e, "long")} 
                  >
                  <i className="fa fa-pencil-square-o" aria-hidden="true"></i> Post Long Question
                </button> 					 					
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
}
