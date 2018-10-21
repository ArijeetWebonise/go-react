import _ from "lodash";

function PING(state = {
    pong: false,
    pin: false,
    res: null,
    error: null
}, action) {
    switch(action.type){
        case 'PING_SEND':{
            let newState = _.assign({}, state);
            newState.pin = true;
            return newState;
        }
        case 'PONG':{
            let newState = _.assign({}, state);
            newState.pin = false;
            newState.pong = true;
            newState.res = action.payload;
            return newState;
        }
        case 'PING_PONG_ERROR':{
            let newState = _.assign({}, state);
            newState.pin = false;
            newState.pong = false;
            newState.error = action.payload;
            return newState;
        }
        default:{
            return state;
        }
    }
}

export default PING;
