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

export const getDashboard = async (token, secret) => {
  const result = await axios.get(
    url + "/performance?key=" + secret,
    authenticate(token)
  );
  return result.data;
};
