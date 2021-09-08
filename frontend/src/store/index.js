import { combineReducers, createStore, applyMiddleware } from "redux";
import ReduxPromise from "redux-promise";
import { composeWithDevTools } from "redux-devtools-extension";

import authReducer from "./reducers/auth";
import dashboardReducer from "./reducers/dashboard";
// import websocket from "./middleware/socket";

const rootReducer = combineReducers({
  auth: authReducer,
  dashboard: dashboardReducer,
});

const store = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(ReduxPromise))
);

export default store;
