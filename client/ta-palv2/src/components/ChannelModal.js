import React from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, FormGroup,Label,Input } from 'reactstrap';

export default class ChannelModal extends React.Component {
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
  handleButton(event, priv){
    event.preventDefault();
    if (this.props.mode === "create") 
      this.props.buttonCallback(this.state.name)
    else 
      this.props.buttonCallback(this.props.channelID, this.state.channelName, this.state.channelDesc)
    this.setState({channelName:''})
    this.setState({
        modal: !this.state.modal
      });
  }

  render() {
    return (
      <span>
        <Button className="float-right" onClick={this.toggle}> {this.props.buttonName}</Button>
        <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
          <ModalHeader toggle={this.toggle}>Create Office Hours</ModalHeader>
          <ModalBody>
          <FormGroup>
            <Label for="name">Name</Label>
            <Input onChange = {e => this.handleChange(e)} id="name" 
              type="name" 
              name="name"
              />
          </FormGroup>
          </ModalBody>
         
            {this.props.mode === "create" ?
             <ModalFooter>  
              <Button color="primary" onClick={(e) => this.handleButton(e)}>Create Office Hour</Button>
              <Button color="secondary" onClick={this.toggle}>Cancel</Button>
              </ModalFooter>
            :
            <ModalFooter>  
              <Button color="primary" onClick={(e) => this.handleButton(e, false)}>Edit Channel</Button>
              <Button color="secondary" onClick={this.toggle}>Cancel</Button>
            </ModalFooter>
            }
            
          
        </Modal>
      </span>
    );
  }
}