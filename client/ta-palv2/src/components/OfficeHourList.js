import React, { Component } from 'react'; //import React Component
import { ListGroup, ListGroupItem} from 'reactstrap'
import { Redirect } from 'react-router-dom'
import Websocket from 'react-websocket';
import "./styles/Questions.css";



export default class OfficeHourList extends Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
    this.state = {
      isOpen: false
    };
  }
  toggle() {
    this.setState({
      isOpen: !this.state.isOpen
    });
  }
  componentDidMount() {
    var auth = localStorage.getItem('Authorization');
    this.setState({user : JSON.parse(localStorage.getItem('User'))})
    fetch("https://tapalapi.patrickold.me/v1/officehours", {
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
          this.setState({ officeHours: response });
        })
        .catch(function(error) {
            localStorage.removeItem("Authorization")
            error.text().then(error => alert("error"))
        })
  }

  update(){
    var auth = localStorage.getItem('Authorization');
    this.setState({user : JSON.parse(localStorage.getItem('User'))})
    fetch("https://tapalapi.patrickold.me/v1/officehours", {
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
          this.setState({ officeHours: response });
        })
        .catch(function(error) {
            error.text().then(error => alert("error"))
        })
  }

  componentWillUnmount() {
    //this.channelRef.off();
  }


  officeHourClickHandler(id) {
    console.log(id)
    this.setState({ clickedOfficeHour: id, redirect: true });
  }

  handleData(data) {
    var message = Buffer.from(data, 'base64').toString('ascii')
    console.log(message)
    this.update();
    if (message === "question-yourTurn") {
      alert("your question is being answered!")
    }
  }

  render() {
    let officeHoursItems = [];
    let officeHours = this.state.officeHours;
    if (officeHours) {
      let keyArray = Object.keys(officeHours).map(key => {
        officeHours[key].dbID =  officeHours[key].id
        officeHours[key].id = key;
        return officeHours[key];
      });
      // keyArray.sort((a, b) => b.time - a.time);
      officeHoursItems = keyArray.map(each => <OfficeHourItem key={each.id} 
        user={this.state.user} deleteOfficeHourCallback={this.props.deleteOfficeHourCallback} 
        officeHourClicked={(id) => this.officeHourClickHandler(id)} officeHour={each} />)
    }

    return (this.state.redirect ?
      (<Redirect to={"/officeHour/" + this.state.clickedOfficeHour} push/>) :
      (
        <div className="container">
          <ListGroup aria-live="polite">
            {officeHoursItems}
          </ListGroup>
          <Websocket url={'wss://tapalapi.patrickold.me/v1/ws?auth=' + localStorage.getItem('Authorization')}
              onMessage={this.handleData.bind(this)}/>
        </div>))
  }
}

class OfficeHourItem extends Component {

  handleClick(officeHour) {
    this.props.officeHourClicked(officeHour.dbID)
  }

  deleteChannelHandler(event) {
    this.props.deleteChannelCallback(this.props.officeHour.dbID)
  }

  render() {
    return (
      <ListGroupItem>
        <span id="oh-name" onClick={(e) => this.handleClick(this.props.officeHour)}>
        {this.props.officeHour.name + "   "}
        </span>
      </ListGroupItem>
    );
  }
}
