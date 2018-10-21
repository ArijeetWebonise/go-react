import React from 'react';
import { HashRouter as Router, Route, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';
import Home from './container/Home';
import Ping from './container/Ping';
import store from './store';

const App = () => (
  <Provider store={store}>
    <Router>
      <Switch>
        <Route
          exact
          path="/"
          component={Home}
          />
        <Route
          path="/ping"
          component={Ping}
          />
      </Switch>
    </Router>
  </Provider>
);

export default App;
