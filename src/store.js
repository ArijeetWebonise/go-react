import { createStore, applyMiddleware } from 'redux';
import reducer from './reducer';
import { createLogger } from 'redux-logger';
import thunk from 'redux-thunk';
 
const logger = createLogger();

const middleware = applyMiddleware(logger, thunk);

export default createStore(reducer, middleware);
