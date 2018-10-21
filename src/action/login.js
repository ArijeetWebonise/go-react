import axios from 'axios';

export function LoginResponse(user) {
  return {
    type: 'USER_LOGGEDIN',
    payload: user,
  };
}

export function LoginFailed(err) {
  return {
    type: 'USER_LOGIN_ERROR',
    payload: err,
  };
}

export function SetLoginFormUser(data) {
  return {
    type: 'USER_FORM_SET_USER',
    payload: data,
  };
}

export function SetLoginFormPass(data) {
  return {
    type: 'USER_FORM_SET_PASS',
    payload: data,
  };
}

export function Login(dispatch, form = { user: '', pass: '' }) {
  const csrfElement = document.getElementById('crsf');
  const ajaxRequest = axios.create({
    headers: { 'X-CSRF-Token': csrfElement.children[0].value },
  });

  const data = {
    user: form.user,
    pass: form.pass,
  };

  ajaxRequest.post('/api/v1/login', data)
    .then((res) => {
      dispatch(LoginResponse(res.data));
    })
    .catch((e) => {
      dispatch(LoginFailed(e));
    });

  return {
    type: 'USER_LOGGING',
  };
}
