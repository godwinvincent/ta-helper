import React from 'react';
import { Button, Modal, ModalHeader, ModalBody, ModalFooter, FormGroup,Label,Input } from 'reactstrap';

export default class QuestionModal extends React.Component {
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
  handleButton(event){
    event.preventDefault();
    this.props.buttonCallback(this.props.question.dbID, this.state.question)
    this.setState({question:''})
    this.setState({
        modal: !this.state.modal
      });
  }

  render() {
    return (
      <span>
        <Button className="float-right" onClick={this.toggle}> Edit Question</Button>
        <Modal isOpen={this.state.modal} toggle={this.toggle} className={this.props.className}>
          <ModalHeader toggle={this.toggle}>Edit Question</ModalHeader>
          <ModalBody>
          <FormGroup>
            <Label for="name">Question</Label>
            <Input onChange = {e => this.handleChange(e)} id="question" 
              type="question" 
              name="question"
              />
          </FormGroup>
          </ModalBody>
            <ModalFooter>  
              <Button color="primary" onClick={(e) => this.handleButton(e, false)}>Edit Question</Button>
              <Button color="secondary" onClick={this.toggle}>Cancel</Button>
            </ModalFooter>
                      
        </Modal>
      </span>
    );
  }
}