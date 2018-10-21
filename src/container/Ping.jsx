import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Ping } from '../action/ping';

class Login extends Component {
  constructor(props) {
    super(props);
    props.Ping();
  }

  render() {
    const ping = this.props.ping.pong ? "ping" : "pong";
    return (
      <div>{ping}</div>
    );
  }
}

const mapStateToProp = state => ({
  ping: state.ping,
});

const mapDispatchToProp = dispatch => ({
  Ping: () => {
    dispatch(Ping(dispatch));
  },
});

export default connect(mapStateToProp, mapDispatchToProp)(Login);
