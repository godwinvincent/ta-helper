import React from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, FormGroup,Label,Input } from 'reactstrap';

export default class UserChanelModal extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      modal: false
    };

    this.toggle = this.toggle.bind(this);
  }

  toggle() {
    this.setState({
      modal: !this.state.modal
    });
  }
  handleChange(event){
    let newState = {};
    newState[event.target.name] = event.target.value;
    this.setState(newState);
  }
  handleButton(event, add){
    event.preventDefault();
    this.props.buttonCallback(this.props.channelID, this.state.userID, add)
    this.setState({userID:''})
    this.setState({
        modal: !this.state.modal
      });
  }

  render() {
    return (
      <span>
        <Button className="float-right" onClick={this.toggle}> Add/Remove User</Button>
        <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
          <ModalHeader toggle={this.toggle}>Channel</ModalHeader>
          <ModalBody>
          <FormGroup>
            <Label for="name">User ID </Label>
            <Input onChange = {e => this.handleChange(e)} id="userID" 
              type="userID" 
              name="userID"
              />
          </FormGroup>
          </ModalBody>
            <ModalFooter>  
              <Button color="primary" onClick={(e) => this.handleButton(e, true)}>Add</Button>
              <Button color="primary" onClick={(e) => this.handleButton(e, false)}>Remove</Button>
              <Button color="secondary" onClick={this.toggle}>Cancel</Button>
            </ModalFooter>
        </Modal>
      </span>
    );
  }
}