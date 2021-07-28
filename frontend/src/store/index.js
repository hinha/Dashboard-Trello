import { createStore, combineReducers } from "redux";
import authReducer from "./reducers/auth";
import usersReducer from "./reducers/socket";
import dashboardReducer from "./reducers/dashboard";

const rootReducer = combineReducers({
  auth: authReducer,
  socket: usersReducer,
  dashboard: dashboardReducer,
});

const store = createStore(rootReducer);

export default store;
