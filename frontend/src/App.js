import React from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import PublicRoute from "./routes/PublicRoute";
import PrivateRoute from "./routes/PrivateRoute";

import Login from "./modules/login/Login";

import Main from "./modules/main/Main";

function App() {
  return (
    <>
      <Router>
        <Switch>
          <PublicRoute exact path="/login">
            <Login />
          </PublicRoute>
          <PrivateRoute path="/dashboard">
            <Main />
          </PrivateRoute>
        </Switch>
      </Router>
    </>
  );
}

export default App;
