import { combineReducers } from 'redux';

// Redux Reducers

const loggedReducer = (state = false, action) => {
  switch (action.type) {
    case 'SIGN_IN':
      return true;
    case 'SIGN_OUT':
      return false;
    default:
      return state;
  }
}

// Combine and export all reducers
const reducers = combineReducers({
  loggedIn: loggedReducer,
});

export default reducers;
