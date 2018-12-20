import React, { Component } from 'react';
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
} from 'reactstrap';
import PropTypes from 'prop-types';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../../theme/creative.sass';

class NavBarComponent extends Component {
  constructor(props) {
    super(props);
    this.toggle = this.toggle.bind(this);
    this.state = {
      isOpen: false,
      navFixed: window.pageYOffset >= 100,
    };
  }

  componentDidMount() {
    window.addEventListener('scroll', () => {
      this.setState({ navFixed: window.pageYOffset >= 100 });
    });
  }

  toggle() {
    const { isOpen } = this.state;
    this.setState({
      isOpen: !isOpen,
    });
  }

  goto(url) {
    const { history } = this.props;
    history.push(url);
  }

  render() {
    const { isOpen, navFixed } = this.state;

    let navClass = 'fixed-top';
    if (window.location.pathname !== '/') {
      navClass = 'navbar-shrink';
    } else if (navFixed) {
      navClass = 'navbar-shrink fixed-top';
    }

    return (
      <Navbar
        id="mainNav"
        color="light"
        light
        fixed={(navFixed) ? 'top' : ''}
        expand="md"
        className={navClass}
        >
        <NavbarBrand
          href="javascript:void(0);"
          onClick={() => this.goto('/')}
          >
          reactstrap
        </NavbarBrand>
        <NavbarToggler onClick={this.toggle} />
        <Collapse
          isOpen={isOpen}
          navbar
          >
          <Nav
            className="ml-auto"
            navbar
            >
            <NavItem>
              <NavLink
                href="javascript:void(0);"
                onClick={() => this.goto('/login')}
                >
                Login
              </NavLink>
            </NavItem>
          </Nav>
        </Collapse>
      </Navbar>
    );
  }
}
NavBarComponent.propTypes = {
  history: PropTypes.shape().isRequired,
};

export default NavBarComponent;
