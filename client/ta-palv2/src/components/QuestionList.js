import React, { Component } from 'react'; //import React Component
import Time from 'react-time'
import './styles/Questions.css';
import QuestionModal from './QuestionModal';
import {Button} from 'reactstrap'


export default class QuestionList extends Component {
  constructor(props){
    super(props);
    this.state = {};
  }

  componentWillMount(){
    console.log(this.props.questions)
  }

  render() {
    let questionItems = [];
    let questions = this.props.questions;
    if(questions){
      let keyArray = Object.keys(questions).map(key => {
        questions[key].dbID = questions[key].id 
        questions[key].id = key;
        return questions[key];
      });
      keyArray.sort((a,b) => a.questPos - b.questPos);
      questionItems = keyArray.map( each =>  <QuestionItem  key={each.id} 
        deleteQuestionCallback={this.props.deleteQuestionCallback} 
        editQuestionCallback={this.props.editQuestionCallback} question={each}
        changeQuestionOrder={this.props.changeQuestionOrder} 
        changeQuestionUsers={this.props.changeQuestionUsers}
        currentUser={this.props.currentUser}/>)
    }
    return (
      <div aria-live="polite" className="container">
          {questionItems}
      </div>);
  }
}


class QuestionItem extends Component {
  deleteQuestionHandler() {
    this.props.deleteQuestionCallback(this.props.question.dbID)
  }

  changeQuestionOrder(change) {
    this.props.changeQuestionOrder(change, this.props.question.dbID)
  }

  changeQuestionUsers(operation) {
    this.props.changeQuestionUsers(this.props.question.dbID, operation)
  }

  sendNotification() {
    this.props.sendNotification(this.props.question.students)
  }




  render() {
    let question = this.props.question;
    let user = JSON.parse(localStorage.getItem('User'));
    console.log(question)
    return (
      <div className="row py-4 bg-white border">
        <div className="col pl-4 pl-lg-1">
          <div id='arrows'>
           {user.role == "instructor" ?
           <span>
            <img className='arrow-buttons' src={window.location.origin + '/img/bell.jpg'} onClick={() => this.sendNotification()}></img>
            <img className='arrow-buttons' src={window.location.origin + '/img/down-arrow.png'} onClick={() => this.changeQuestionOrder('down')}></img>
            <img className='arrow-buttons' src={window.location.origin + '/img/up-arrow.png'} onClick={() => this.changeQuestionOrder('up')}></img>
            <img className='arrow-buttons' src={window.location.origin + '/img/minus.svg'} onClick={() => this.changeQuestionUsers('remove')}></img>
            </span>
           : 
           <span>
             {question.students.includes(user.username) ? 
             <img className='arrow-buttons' src={window.location.origin + '/img/minus.svg'} onClick={() => this.changeQuestionUsers('remove')}></img>
            :
            <img className='arrow-buttons' src={window.location.origin + '/img/plus.jpg'} onClick={() => this.changeQuestionUsers('add')}></img> 
            }
            </span>
            }
          </div>
          <div className="question">{question.questBody}</div>
          <div className="Student">{ "Student Count: " + question.students.length}</div>
        </div>
      </div>
    );
  }
}
