import _ from 'lodash';

const initState = {
  loggedin: false,
  logging: false,
  user: null,
  form: { user: '', pass: '' },
  error: null,
};

const login = (state = initState, action) => {
  switch (action.type) {
    case 'USER_LOGGING': {
      const newState = _.assign({}, state);
      newState.logging = true;
      return newState;
    }
    case 'USER_LOGGEDIN': {
      const newState = _.assign({}, state);
      newState.logging = false;
      newState.loggedin = true;
      newState.user = action.payload;
      return newState;
    }
    case 'USER_LOGIN_ERROR': {
      const newState = _.assign({}, state);
      newState.logging = false;
      newState.loggedin = false;
      newState.error = action.payload;
      return newState;
    }
    case 'USER_FORM_SET_USER': {
      const newState = _.assign({}, state);
      const newForm = _.assign({}, newState.form);
      newForm.user = action.payload;
      newState.form = newForm;
      return newState;
    }
    case 'USER_FORM_SET_PASS': {
      const newState = _.assign({}, state);
      const newForm = _.assign({}, newState.form);
      newForm.pass = action.payload;
      newState.form = newForm;
      return newState;
    }
    default: {
      return state;
    }
  }
};

export default login;
