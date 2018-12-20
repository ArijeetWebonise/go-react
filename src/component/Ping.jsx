import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Ping } from '../action/ping';

class Login extends Component {
  constructor(props) {
    super(props);
    props.Ping();
  }

  render() {
    const { ping } = this.props;
    const pingResponse = ping.pong ? 'ping' : 'pong';
    return (
      <div>{pingResponse}</div>
    );
  }
}

Login.propTypes = {
  ping: PropTypes.shape({
    pong: PropTypes.bool.isRequired,
  }).isRequired,
  Ping: PropTypes.func.isRequired,
};

const mapStateToProp = state => ({
  ping: state.ping,
});

const mapDispatchToProp = dispatch => ({
  Ping: () => {
    dispatch(Ping(dispatch));
  },
});

export default connect(mapStateToProp, mapDispatchToProp)(Login);
