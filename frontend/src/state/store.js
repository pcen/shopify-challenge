import { createStore, applyMiddleware, compose } from 'redux';
import thunkMiddleware from 'redux-thunk';
import reducers from './reducers';

// Configure Redux store
// - apply thunk middleware
// - enable redux devtools
const middleware = applyMiddleware(thunkMiddleware);
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const enhancers = composeEnhancers(middleware);
const store = createStore(reducers, enhancers);

export default store;
