import * as ActionTypes from "../actions";

const initialState = {
  isLoggedIn: !!localStorage.getItem("token"),
  token: localStorage.getItem("token"),
  credentials: localStorage.getItem("credential"),
  currentUser: {
    name: "user",
    username: "user",
    email: "mail@example.com",
    picture: null,
  },
  isArn: !!localStorage.getItem("arn"),
  arnData: [],
};

const reducer = (state = initialState, action) => {
  if (action.type === ActionTypes.LOGIN_USER) {
    localStorage.setItem("token", action.token);
    return {
      ...state,
      isLoggedIn: true,
      token: action.token,
    };
  }

  if (action.type === ActionTypes.LOGOUT_USER) {
    localStorage.removeItem("token");
    localStorage.removeItem("performance");
    localStorage.removeItem("arn");
    localStorage.removeItem("credential");
    return {
      ...state,
      isLoggedIn: false,
      token: null,
      currentUser: {
        email: "mail@example.com",
        picture: null,
      },
      credentials: null,
    };
  }
  if (action.type === ActionTypes.LOAD_USER) {
    return {
      ...state,
      currentUser: action.currentUser,
      credentials: action.credentials,
    };
  }

  if (action.type === ActionTypes.ARN_USER) {
    localStorage.setItem("arn", JSON.stringify(action.arnList));
    return {
      ...state,
      isArn: true,
      arnData: action.arnList,
    };
  }

  if (action.type == ActionTypes.ADD_CREDENTIALS) {
    localStorage.setItem("credential", action.credentials);
    return {
      ...state,
      credentials: action.credentials,
    };
  }

  if (action.type === ActionTypes.LOAD_TOKEN) {
    return {
      ...state,
      token: state.token,
    };
  }

  return { ...state };
};

export default reducer;
