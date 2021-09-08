import * as ActionTypes from "../actions";

const initialState = {
  performance: {},
  settings: {
    accounts: [],
  },
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

  if (action.type === ActionTypes.ADD_USER_FORM) {
    return {
      ...state,
      settings: {
        accounts: action.account,
      },
    };
  }
  if (action.type === ActionTypes.LOAD_USER_FORM) {
    return {
      ...state,
      settings: action.settings,
    };
  }

  return { ...state };
};

export default reducer;
