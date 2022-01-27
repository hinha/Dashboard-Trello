import axios from "axios";

const apiBase = process.env.REACT_APP_API_BASE;
const apiAnalytic = process.env.REACT_APP_URI_API_DASHBOARD;

const authenticate = (token) => ({
  headers: {
    "Content-type": "application/json",
    Authorization: `Bearer ${token}`,
  },
});
const url = `${apiBase}${apiAnalytic}`;

export const getTrelloCard = async (token, credentials) => {
  const result = await axios.get(
    `${url}/analytic/trello?key=${credentials}`,
    authenticate(token)
  );
  return result.data;
};

export const getClusters = async (token, credentials, params = "2020") => {
  const result = await axios.get(
    `${url}/analytic/trello/clusters?key=${credentials}&year=${params}`,
    authenticate(token)
  );
  return result.data;
};

export const getOverviewData = async (token, credentials, params = "2020") => {
  const result = await axios.get(
    `${url}/analytic/trello/data?key=${credentials}&year=${params}`,
    authenticate(token)
  );
  return result.data;
};

export const UPDATE_ANALYTIC_TRELLO = "UPDATE_ANALYTIC_TRELLO";
