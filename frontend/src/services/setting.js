import axios from "axios";

const apiBase = process.env.REACT_APP_API_BASE;
const apiDashboard = process.env.REACT_APP_URI_API_DASHBOARD;

const authenticate = (token) => ({
  headers: {
    "Content-type": "application/json",
    Authorization: `Bearer ${token}`,
  },
});
const url = `${apiBase}${apiDashboard}`;

export const getSettingUser = async (token, credentials) => {
  const result = await axios.get(
    `${url}/settings/user?key=${credentials}`,
    authenticate(token)
  );
  return result.data;
};

export const roleSettingUser = async (token, credentials, body) => {
  try {
    const result = await axios.post(
      `${url}/settings/user/role?key=${credentials}`,
      body,
      authenticate(token)
    );
    return result.data;
  } catch (e) {
    return e.response.data;
  }
};

export const addSettingUser = async (token, credentials, body) => {
  try {
    const result = await axios.post(
      `${url}/settings/user?key=${credentials}`,
      body,
      authenticate(token)
    );
    return result.data;
  } catch (e) {
    return e.response.data;
  }
};

export const editSettingUser = async (token, credentials, body) => {
  try {
    const result = await axios.patch(
      `${url}/settings/user?key=${credentials}`,
      body,
      authenticate(token)
    );
    return result.data;
  } catch (e) {
    return e.response.data;
  }
};

export const delSettingUser = async (token, credentials, params = "") => {
  try {
    const result = await axios.delete(
      `${url}/settings/user?key=${credentials}&s=${params}`,
      authenticate(token)
    );
    return result.data;
  } catch (e) {
    return e.response.data;
  }
};

export const trelloSettingUser = async (token, credentials, body) => {
  try {
    const result = await axios.post(
      `${url}/settings/user/trello?key=${credentials}`,
      body,
      authenticate(token)
    );
    return result.data;
  } catch (e) {
    return e.response.data;
  }
};

export const UPDATE_USER_SETTING = "UPDATE_USER_SETTING";
export const ADD_USER_SETTING = "ADD_USER_SETTING";
export const DEL_USER_SETTING = "DELETE_USER_SETTING";
export const EDIT_USER_SETTING = "EDIT_USER_SETTING";
export const TRELLO_USER_SETTING = "TRELLO_USER_SETTING";
export const ROLE_USER_SETTING = "ROLE_USER_SETTING";
