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
        {userPull.role === "instructor" ?
          <NavItem>
            <NavLink>
              <ChannelModal mode="create" buttonName="New Office Hours" buttonCallback={this.props.newOfficeHourCallback} />
            </NavLink>
          </NavItem> : '' }
          <NavItem>
            <NavLink>
              <Button className="btn btn-warning" onClick={() => this.handleSignOut()}> Log Out </Button>
            </NavLink>
          </NavItem>
        </Nav>
      </Collapse>)
    }

    return (
      <div>
        <Navbar color="faded" light expand="md">
          <Link to='/'>
            <NavbarBrand>
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