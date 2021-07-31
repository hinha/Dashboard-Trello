import { applyMiddleware, compose, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";
import websocket from "@giantmachines/redux-websocket";
import { getDefaultMiddleware } from "@reduxjs/toolkit";

import authReducer from "./reducers/auth";
import usersReducer from "./reducers/socket";
import dashboardReducer from "./reducers/dashboard";
// import websocket from "./middleware/socket";

const websocketMiddleware = websocket({
  prefix: "REDUX_WEBSOCKET",
  onOpen: (socket) => {
    // @ts-ignore
    window.__socket = socket; // eslint-disable-line no-underscore-dangle
  },
  dateSerializer: (date) => date.toISOString(),
  deserializer: JSON.parse,
});

/**
 * Create a custom middleware to make the simulate disconnect feature function
 * as a more interesting demo. Force the middleware to attempt to reconnect
 * an arbitrary number of times.
 */
const disconnectSimulatorMiddleware = () => {
  const OldWebSocket = window.WebSocket;
  return (next) => (action) => {
    const { type, payload } = action;
    // If the connection breaks, block reconnection by overwriting the WebSocket
    // class with a fake class.
    if (type === `REDUX_WEBSOCKET::BROKEN`) {
      window.WebSocket = class FakeWebSocket {
        constructor() {
          this.close = () => {};
          this.addEventListener = () => {};
        }
      };
    }
    // Monitor how many reconnnection attempts were made, and if we had
    // enough, allow a reconnect to happen by restoring the original WebSocket.
    if (type === `REDUX_WEBSOCKET::RECONNECT_ATTEMPT`) {
      const { count } = payload;
      if (count > 2) {
        window.WebSocket = OldWebSocket;
      }
    }

    return next(action);
  };
};

const rootReducer = combineReducers({
  auth: authReducer,
  socket: usersReducer,
  dashboard: dashboardReducer,
});

const store = createStore(
  rootReducer,
  compose(
    applyMiddleware(
      thunk,
      disconnectSimulatorMiddleware,
      websocketMiddleware,
      ...getDefaultMiddleware({
        serializableCheck: {
          ignoredActionPaths: ["payload.event"],
          ignoredActions: ["REDUX_WEBSOCKET::OPEN"],
        },
      })
    ),
    window.devToolsExtension ? window.devToolsExtension() : (f) => f
  )
);

// const immutableInvariantMiddleware = createImmutableStateInvariantMiddleware({
//   ignoredActionPaths: ["payload.event"],
//   ignoredActions: ["REDUX_WEBSOCKET::OPEN"],
// });

// const store = configureStore({
//   reducer: rootReducer,
//   middleware: [
//     disconnectSimulatorMiddleware,
//     websocketMiddleware,
//     immutableInvariantMiddleware,
//     thunk,
//   ],
// });

export default store;
