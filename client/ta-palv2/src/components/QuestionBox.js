import React, { Component } from 'react'; //import React Component
import './styles/Questions.css';

//A form the user can use to post a Chirp
export default class QuestionBox extends Component {
  constructor(props){
    super(props);
    this.state = {question:''};
  }

  post(body) {
    var questionData = {
        "body": body
    }
    var auth = localStorage.getItem('Authorization');
    fetch("http://localhost:80/v1/questions/{officeHourID}" + this.props.id, {
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
            error.text().then(error => this.setState({ errorMessage: error }))
        })
  }

  updateQuestion(event) {
    console.log("updating")
    this.setState({question: event.target.value});
  }
  postQuestion(event){
    event.preventDefault(); 
    this.post(this.state.question)
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
                  onClick={(e) => this.postQuestion(e)} 
                  >
                  <i className="fa fa-pencil-square-o" aria-hidden="true"></i> Send
                </button> 					
              </div>
            </form>
          </div>
        </div>
      </div>
    );
  }
}
