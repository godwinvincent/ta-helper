import React, { Component } from 'react'; //import React Component
import './styles/Messages.css';

//A form the user can use to post a Chirp
export default class MessageBox extends Component {
  constructor(props){
    super(props);
    this.state = {message:''};
  }

  post(body) {
    var messageData = {
        "body": body
    }
    var auth = localStorage.getItem('Authorization');
    fetch("https://info441api.godwinv.com/v1/channels/" + this.props.id, {
            method: "POST",
            mode: "cors", 
            headers: {
                "Content-Type": "application/json",
                "Authorization": auth
            },
            body: JSON.stringify(messageData), 
        })
        .then(response => {
            if (response.status < 300) {
                this.setState({message:''}); 
            } else {
                throw response
            }
        })
        .catch(function(error) {
            error.text().then(error => this.setState({ errorMessage: error }))
        })
  }

  updateMessage(event) {
    console.log("updating")
    this.setState({message: event.target.value});
  }
  postMessage(event){
    event.preventDefault(); 
    this.post(this.state.message)
    this.setState({message:''}); 
  }


  render() {
    return (
      <div className="container">
        <div className="row py-3 chirp-box">
          <div className="col pl-4 pl-lg-1">
            <form>
              <textarea name="text" className="form-control mb-2" placeholder="Type Message Here" 
                value={this.state.message} 
                onChange={(e) => this.updateMessage(e)}
                />
              
              <div className="text-right">
                <button className="btn btn-primary" 
                  onClick={(e) => this.postMessage(e)} 
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
