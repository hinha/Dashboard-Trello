import * as ActionTypes from "../actions";

const initialState = {
  socket: null,
};

const reducer = (state = initialState, action) => {
  if (action.type === ActionTypes.ADD_SOCKET) {
    return {
      ...state,
      socket: action.socket,
    };
  }

  if (action.type === ActionTypes.LOAD_SOCKET) {
    return {
      ...state,
      socket: action.socket,
    };
  }

  return { ...state };
};

export default reducer;
