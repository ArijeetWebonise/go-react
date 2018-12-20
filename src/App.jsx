import React, { Fragment } from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';
import { createBrowserHistory } from 'history';
import Home from './component/Home/home';
import Ping from './component/Ping';
import Login from './component/Login/Login';
import store from './store';
import Navbar from './component/NavBar/Navbar';

const history = createBrowserHistory();

const App = () => (
  <Fragment>
    <Navbar history={history} />
    <Provider store={store}>
      <Router history={history}>
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
          <Route
            path="/login"
            component={Login}
            />
        </Switch>
      </Router>
    </Provider>
  </Fragment>
);

export default App;
