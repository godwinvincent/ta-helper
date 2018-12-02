import React, { Component } from 'react';
//import './Login.css'
import Popup from 'reactjs-popup'

class Question extends Component {
    constructor(props) {
      super(props);
      this.state = {
        value: null,
      };
    }
  
    render() {
      return (
            <Popup
              trigger={<button className="button"> Open Modal </button>}
              modal
              closeOnDocumentClick
            >
              <span> Modal content </span>
            </Popup>
          )

    }
  }

  export default Question;