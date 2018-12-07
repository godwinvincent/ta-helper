import React from 'react';
import { Collapse, Navbar, NavbarToggler, NavbarBrand, Nav, NavItem, NavLink, Button } from 'reactstrap';
import ChannelModal from './ChannelModal'
import {Link} from 'react-router-dom'

export default class Header extends React.Component {
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

  handleSignOut() {
    this.props.signOutCallback();
  }

  render() {
    let content = "";
    var userPull = JSON.parse(localStorage.getItem("User"))
    if (this.props.showOptions) {
      content = (<Collapse isOpen={this.state.isOpen} navbar>
        <Nav className="ml-auto" navbar>
        {userPull.role === "instructor" && !this.props.onlyLogout ?
          <NavItem>
            <NavLink>
              <ChannelModal mode="create" buttonName="New Office Hours" buttonCallback={this.props.newOfficeHourCallback} />
            </NavLink>
          </NavItem> : '' }
          <NavItem>
            <NavLink>
              <Button className="btn btn-danger" onClick={() => this.handleSignOut()}> Log Out </Button>
            </NavLink>
          </NavItem>
        </Nav>
      </Collapse>)
    }

    let styles = {
      background: "#4b2e83",
      marginBottom: "100px"
    };

    return (
      <div style = {styles}>
        <Navbar color="faded" light expand="md">
          <Link to='/'>
            <NavbarBrand className="text-white">
              TA-PAL
            </NavbarBrand>
          </Link>
          <NavbarToggler onClick={this.toggle} />
          {content}
        </Navbar>
      </div>
    );
  }
}