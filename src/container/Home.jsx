import React, { Component } from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { Login, SetLoginFormUser, SetLoginFormPass } from '../action/login';

class Home extends Component {
  constructor(props) {
    super(props);
    this.handleLogin = this.handleLogin.bind(this);
    this.handleChangeUser = this.handleChangeUser.bind(this);
    this.handleChangePass = this.handleChangePass.bind(this);
  }

  handleLogin() {
    this.props.Login(this.props.login.form.user, this.props.login.form.pass);
  }

  handleChangeUser(e) {
    this.props.SetLoginFormUser(e.currentTarget.value);
  }

  handleChangePass(e) {
    this.props.SetLoginFormPass(e.currentTarget.value);
  }

  render() {
    return (
      <div>
        <h1>Home</h1>
        <div>
          <input type="text" name="username" onChange={this.handleChangeUser} />
        </div>
        <div>
          <input type="text" name="password" onChange={this.handleChangePass} />
        </div>
        <button onClick={this.handleLogin} type="button">Login</button>
      </div>
    );
  }
}

Home.protoTypes = {
  login: PropTypes.shape({
    user: PropTypes.any,
    form: PropTypes.shape({
      user: PropTypes.string,
      pass: PropTypes.string,
    }),
  }).isRequired,
};

const mapStateToProp = state => ({
  login: state.login,
});

const mapDispatchToProp = dispatch => ({
  Login: (username, password) => {
    dispatch(Login(dispatch, { user: username, pass: password }));
  },
  SetLoginFormUser: (data) => {
    dispatch(SetLoginFormUser(data));
  },
  SetLoginFormPass: (data) => {
    dispatch(SetLoginFormPass(data));
  },
});

export default connect(mapStateToProp, mapDispatchToProp)(Home);
