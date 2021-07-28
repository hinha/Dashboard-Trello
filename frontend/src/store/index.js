import { createStore, combineReducers } from "redux";
import authReducer from "./reducers/auth";
import usersReducer from "./reducers/socket";

const rootReducer = combineReducers({
  auth: authReducer,
  socket: usersReducer,
});

const store = createStore(rootReducer);

export default store;
