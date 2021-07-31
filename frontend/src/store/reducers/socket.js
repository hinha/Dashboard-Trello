import {
  REDUX_WEBSOCKET_BROKEN,
  REDUX_WEBSOCKET_CLOSED,
  REDUX_WEBSOCKET_CONNECT,
  REDUX_WEBSOCKET_ERROR,
  REDUX_WEBSOCKET_MESSAGE,
  REDUX_WEBSOCKET_OPEN,
  REDUX_WEBSOCKET_SEND,
} from "../actions";

const defaultState = {
  connected: false,
  messages: [],
  url: null,
};

export const getConnected = (state) => state.connected;

const reducer = (state = defaultState, action) => {
  switch (action.type) {
    case "INTERNAL::CLEAR_MESSAGE_LOG":
      return {
        ...state,
        messages: [],
      };
    case REDUX_WEBSOCKET_CONNECT:
      return {
        ...state,
        url: action.payload.url,
      };
    case REDUX_WEBSOCKET_OPEN:
      return {
        ...state,
        connected: true,
      };
    case REDUX_WEBSOCKET_BROKEN:
      return {
        ...state,
        connected: false,
      };
    case REDUX_WEBSOCKET_CLOSED:
      return {
        ...state,
        connected: false,
      };
    case REDUX_WEBSOCKET_MESSAGE:
      return {
        ...state,
        messages: [
          {
            data: action.payload.message,
            origin: action.payload.origin,
            timestamp: action.meta.timestamp,
            type: "INCOMING",
          },
        ],
      };
    case REDUX_WEBSOCKET_SEND:
      return {
        ...state,
        messages: [
          {
            data: action.payload,
            origin: window.location.origin,
            timestamp: new Date().toISOString(),
            type: "OUTGOING",
          },
        ],
      };
    case REDUX_WEBSOCKET_ERROR:
      return {
        ...state,
        connected: true,
      };
    default:
      return state;
  }
};
export default reducer;
