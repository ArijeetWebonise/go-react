import React, { Component } from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { Login } from '../action/login';

class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user: '',
      pass: '',
    };
    this.handleLogin = this.handleLogin.bind(this);
    this.handleChangeUser = this.handleChangeUser.bind(this);
    this.handleChangePass = this.handleChangePass.bind(this);
    this.getMsg = this.getMsg.bind(this);
  }

  getMsg() {
    const { props } = this;
    if (props.login.error !== null) {
      return (<div>Email or Password Failed</div>);
    }
    return null;
  }

  handleLogin() {
    const { props, state } = this;
    props.Login(state.user, state.pass);
  }

  handleChangeUser(e) {
    this.setState({
      user: e.currentTarget.value,
    });
  }

  handleChangePass(e) {
    this.setState({
      pass: e.currentTarget.value,
    });
  }

  render() {
    return (
      <div>
        <h1>Home</h1>
        {this.getMsg()}
        <div>
          <input
            type="text"
            name="username"
            onChange={this.handleChangeUser}
            />
        </div>
        <div>
          <input
            type="text"
            name="password"
            onChange={this.handleChangePass}
            />
        </div>
        <button
          onClick={this.handleLogin}
          type="button"
          >
          <p>Login</p>
        </button>
      </div>
    );
  }
}

Home.propTypes = {
  login: PropTypes.shape({
    user: PropTypes.any,
    form: PropTypes.shape().isRequired,
  }).isRequired,
};

const mapStateToProp = state => ({
  login: state.login,
});

const mapDispatchToProp = dispatch => ({
  Login: (username, password) => {
    dispatch(Login(dispatch, { user: username, pass: password }));
  },
});

export default connect(mapStateToProp, mapDispatchToProp)(Home);
