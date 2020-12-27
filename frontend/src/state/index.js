import { createStore, applyMiddleware, compose, combineReducers } from 'redux';
import thunkMiddleware from 'redux-thunk';

import userReducer from './reducers/user';

// Combine reducers
const reducers = combineReducers({
  user: userReducer,
});

// Configure Redux store
// - apply thunk middleware
// - enable redux devtools
const middleware = applyMiddleware(thunkMiddleware);
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const enhancers = composeEnhancers(middleware);
const store = createStore(reducers, enhancers);

export { store };
