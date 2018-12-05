import React, { Component } from 'react';
import SignUpForm from './components/SignUpForm';
import SignInForm from './components/SignInForm';
import OfficeHour from './components/OfficeHour'
import Home from './components/Home';
import { BrowserRouter as Router, Route} from 'react-router-dom'
import './App.css';


class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      shouldRedirect: false,
    };
  }

  async signup(netid, password, passwordConf, firstName, lastName) {
    var jsonData = {
        "email": netid+"@uw.edu",
        "password": password,
        "passwordConf": passwordConf,
        "userName": netid,
        "firstName": firstName,
        "lastName": lastName
    }
    console.log(jsonData)
    fetch("http://localhost:80/v1/users", {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(jsonData), // body data type must match "Content-Type" header
        })
        .then(response => {
            if (response.status < 300) {
                window.localStorage.setItem('Authorization', response.headers.get("Authorization"));
                return response.json()
            } else {
                throw response
            }
        })
        .then(user => {
          if (user) {
            window.localStorage.setItem('User', JSON.stringify(user));
            this.setState({ user: user, shouldRedirect: true })
          }
          else {
            this.setState({ user: null })
          }
          this.setState({ loading: false });
        })
        .catch(function(error) {
          error.text().then(error => this.setState({ errorMessage: error }))
        })

  }

  async signin(netid, password) {
  var jsonData = {
      "email": netid + "@uw.edu",
      "password": password
  }
  fetch("http://localhost:80/v1/sessions", {
          method: "POST", // *GET, POST, PUT, DELETE, etc.
          mode: "cors", // no-cors, cors, *same-origin
          headers: {
              "Content-Type": "application/json"
          },
          body: JSON.stringify(jsonData), // body data type must match "Content-Type" header
      })
      .then(response => {
          if (response.status < 300) {
              window.localStorage.setItem('Authorization', response.headers.get("Authorization"));
              return response.json()
          } else {
              throw response
          }
      })
      .then(user => {
        if (user) {
          window.localStorage.setItem('User', JSON.stringify(user));
          this.setState({ user: user, shouldRedirect: true })
        }
        else {
          this.setState({ user: null })
        }
        this.setState({ loading: false });
      })
      .catch(function(error) {
          error.text().then(error => alert(error))
      })
  }

  async signout() {
    var auth = localStorage.getItem('Authorization');
    fetch("http://localhost:80/v1/sessions/mine", {
            method: "DELETE", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, cors, *same-origin
            headers: {
                "Authorization": auth
            },
        })
        .then(response => {
            if (response.status < 300) {
                return response.text()
            } else {
                throw response
            }
        })
        .then(response => {
            localStorage.removeItem('Authorization');
            localStorage.removeItem('User');
            this.setState({shouldRedirect: false})
        })
        .catch( (error) => {
            error.text().then(error => this.setState({ errorMessage: error }))
        })


  }

  componentDidMount() {
    var auth = localStorage.getItem('Authorization');
    if (auth) {
      this.setState({ user: auth, shouldRedirect: true })
    }
    else {
      this.setState({ user: null })
    }
    
    this.setState({ loading: false });
  }

  componentWillUnmount() {
  }

  handleSignUp(email, password, passwordConf, username, firstName, lastName) {
    this.setState({ errorMessage: null }); 
    this.signup(email, password, passwordConf, username, firstName, lastName)
  }

  handleSignIn(email, password) {
    this.setState({ errorMessage: null });
    this.signin(email, password)

  }

  handleSignOut() {
    this.setState({ errorMessage: null }); 
    this.signout()
    localStorage.removeItem('Authorization');
    this.setState({shouldRedirect: false, user: null})

  }

  render() {
    let styles = {position:'fixed',left:0,bottom:0,width:'100%'};
    return(
        <div>
          <Router basename={process.env.PUBLIC_URL + '/'}>
          <div>
            <Route exact path="/" render={() => {
              return <Home user={this.state.user} signOutCallback = {() => this.handleSignOut()} loading={this.state.loading} />
            }} />
            <Route exact path="/officeHour" render={() => {
              return <Home user={this.state.user} signOutCallback = {() => this.handleSignOut()} loading={this.state.loading}/>
            }} />
            <Route path="/login" render={(routerProps) => (
              <div className="container">
                <SignInForm {...routerProps}
                  signInCallback={(n, p) => this.handleSignIn(n, p)}
                  redirect={this.state.shouldRedirect}
                />
              </div>
            )} />
            <Route path="/join" render={(routerProps) => (
              <div className="container">
                <SignUpForm {...routerProps}
                  signUpCallback={(n, p, pc, fn, ln) => this.handleSignUp(n, p, pc, fn, ln)}
                  redirect={this.state.shouldRedirect}
                />
              </div>
            )} />
            <Route path="/officeHour/:id" render={(routerProps) => (
              <OfficeHour {...routerProps} currentUser={this.state.user} privateMessage={false} />
            )} push/>
          </div>
        </Router>
        <footer style={styles}>Created by Ben, Godwin, and Patrick</footer>
        </div>
    );
  }
}

export default App;
