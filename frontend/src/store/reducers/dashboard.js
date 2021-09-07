import * as ActionTypes from "../actions";

const initialState = {
  performance: {},
};

const reducer = (state = initialState, action) => {
  if (action.type === ActionTypes.ADD_DATA) {
    return {
      ...state,
      performance: action.performance,
    };
  }

  if (action.type === ActionTypes.LOAD_DATA) {
    return {
      ...state,
      performance: action.performance,
    };
  }

  return { ...state };
};

export default reducer;
