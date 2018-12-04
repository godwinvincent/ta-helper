import React, { Component } from 'react'; //import React Component
import Time from 'react-time'
import './styles/Messages.css';
import MessageModal from './MessageModal';
import {Button} from 'reactstrap'


export default class MessageList extends Component {
  constructor(props){
    super(props);
    this.state = {};
  }

  componentWillMount(){
    console.log(this.props.messages)
  }

  render() {
    let messageItems = [];
    let messages = this.props.messages;
    if(messages){
      let keyArray = Object.keys(messages).map(key => {
        messages[key].dbID = messages[key].id 
        messages[key].id = key;
        return messages[key];
      });
      keyArray.sort((a,b) => a.time - b.time);
      messageItems = keyArray.map( each =>  <MessageItem  key={each.id} deleteMessageCallback={this.props.deleteMessageCallback} editMessageCallback={this.props.editMessageCallback} message={each} currentUser={this.props.currentUser}/>)
    } 

    return (
      <div aria-live="polite" className="container">
          {messageItems}
      </div>);
  }
}


class MessageItem extends Component {
  deleteMessageHandler(){
    this.props.deleteMessageCallback(this.props.message.dbID)
  }
  render() {
    let message = this.props.message;
    let user = JSON.parse(localStorage.getItem('User'));
    console.log(message)
    return (
      <div className="row py-4 bg-white border">
        <div className="col pl-4 pl-lg-1">
          {/* <span className="handle">{message.creator.username} space</span> */}
          {/* <span className="time"><Time value={message.createdAt} relative/></span> */}
          <div className="message">{message.questBody}</div>
        </div>
        {/* {this.props.message.creator.id ===  user.id ?
        <span>
        <MessageModal message={message} buttonCallback={this.props.editMessageCallback}></MessageModal>
        <Button color="danger" className="float-right" onClick={(e) => this.deleteMessageHandler()}>Delete</Button>
        </span> : ''} */}
      </div>      
    );
  }
}
