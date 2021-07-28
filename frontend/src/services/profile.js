import axios from "axios";

const apiBase = process.env.REACT_APP_API_BASE;
const apiLogin = process.env.REACT_APP_URI_API_PROFILE;

const authenticate = (token) => ({
  headers: {
    "Content-type": "application/json",
    Authorization: `Bearer ${token}`,
  },
});
const url = `${apiBase}${apiLogin}`;

export const getProfile = async (token) => {
  const result = await axios.get(url, authenticate(token));
  return result.data;
};

export const refreshToken = async (token) => {
  const result = await axios.post(`${url}/refresh`, {}, authenticate(token));
  return result.data;
};
