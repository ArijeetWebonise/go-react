import ping from './ping';
import login from './login';
import { combineReducers } from 'redux';

const reducers = {
    ping,
    login
}

export default combineReducers(reducers);
