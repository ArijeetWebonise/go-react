import axios from 'axios';

export function Ping(dispatch) {
    let csrfElement = document.getElementById("crsf");
    let ajaxRequest = axios.create({
        headers: {'X-CSRF-Token' : csrfElement.children[0].value}
    });

    ajaxRequest.get('/api/v1/ping')
        .then((res) => {
            dispatch(Pong(res.data));
        })
        .catch(e => {
            dispatch(PingPongError(e));
        });

    return {
        type: 'PING_SEND',
    };
}

export function Pong(data) {
    return {
        type: 'PONG',
        payload: data,
    };
}

export function PingPongError(err) {
    return {
        type: 'PING_PONG_ERROR',
        payload: err,
    };
}
