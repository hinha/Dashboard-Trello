import axios from "axios";
import Cookies from "js-cookie";

const apiBase = process.env.REACT_APP_API_BASE;
const apiLogin = process.env.REACT_APP_URI_API_AUTH;

const configWithCredentials = {
  ...{
    headers: {
      "Content-type": "multipart/form-data",
    },
  },
  withCredentials: true,
};
const url = `${apiBase}${apiLogin}`;

export const getToken = async () => {
  const result = await axios.get(url, configWithCredentials);
  Cookies.set("csrf", result.data);
  return result;
};

export const loginByAuth = async (username, password, remember_me) => {
  let form = new FormData();
  form.set("username", username);
  form.set("password", password);
  form.set("remember", remember_me);

  configWithCredentials.headers["X-CSRF-Token"] = Cookies.get("csrf");
  const result = await axios.post(url, form, configWithCredentials);
  return result;
};
