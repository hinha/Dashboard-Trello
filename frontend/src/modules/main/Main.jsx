import React, { useState, useEffect } from "react";
import { Route, Switch } from "react-router-dom";
import { connect as connectRedux } from "react-redux";
import {
  connect as websocketConnect,
  disconnect,
  send,
} from "@giantmachines/redux-websocket";

import Footer from "./footer/Footer";
import Header from "./header/Header";
import Sidebar from "./../../components/layout/Sidebar";
import PageLoading from "./../../components/page-loading/PageLoading";
import Attendence from "./../../pages/Attendence";
import Home from "./../../pages/Home";
import SettingsDetail from "../../pages/SettingsDetail";
import SettingsUser from "../../pages/SettingsUser";

import * as AuthService from "./../../services/profile";
import * as DashboardService from "./../../services/dashboard";
import * as ActionTypes from "../../store/actions";
import { getConnected } from "../../store/reducers/socket";
// import Socket from "./socket";

const Main = ({
  connected,
  token,
  onCredential,
  onDashboard,
  onUserLoad,
  onUserLogout,
  connect,
}) => {
  const [appLoadingState, updateAppLoading] = useState(false);
  const [menusidebarState, updateMenusidebarState] = useState({
    isMenuSidebarCollapsed: false,
  });

  // const [getToken, setToken] = useState("");

  useEffect(() => {
    updateAppLoading(true);
    let mounted = true;
    const fetchProfile = async () => {
      try {
        const response = await AuthService.getProfile(token);

        onUserLoad({ ...response });
        onCredential(response.credentials);

        const dashboard = await DashboardService.getDashboard(
          token,
          localStorage.getItem("credential")
        );

        if (mounted) {
          onDashboard(dashboard);
          updateAppLoading(false);
        }
      } catch (error) {
        console.log(error);
        if (error.response) {
          if (error.response.status === 401) {
            onUserLogout();
          }
        }
        // updateAppLoading(true);
      }
    };

    fetchProfile();

    if (connected === undefined || connected === false) {
      connect(
        "ws://localhost:8080/dashboard/inbox/ws?key=" +
          localStorage.getItem("credential")
      );
    }

    return () => (mounted = false);
  }, [connected, onUserLoad, onCredential, onDashboard]);

  const toggleMenuSidebar = () => {
    updateMenusidebarState({
      isMenuSidebarCollapsed: !menusidebarState.isMenuSidebarCollapsed,
    });
  };

  document.getElementById("root").classList.remove("register-page");
  document.getElementById("root").classList.remove("login-page");
  document.getElementById("root").classList.remove("hold-transition");

  document.getElementById("root").className += " sidebar-mini";

  if (menusidebarState.isMenuSidebarCollapsed) {
    document.getElementById("root").classList.add("sidebar-collapse");
    document.getElementById("root").classList.remove("sidebar-open");
  } else {
    document.getElementById("root").classList.add("sidebar-open");
    document.getElementById("root").classList.remove("sidebar-collapse");
  }
  let template;

  if (appLoadingState) {
    template = <PageLoading />;
  } else {
    template = (
      <>
        <Header toggleMenuSidebar={toggleMenuSidebar} />
        <Sidebar />
        <div className="content-wrapper">
          <Switch>
            <Route exact path="/dashboard" component={Home} />
            <Route exact path="/dashboard/attendence" component={Attendence} />
            <Route exact path="/analytics" component={Attendence} />
            <Route exact path="/settings" component={SettingsDetail} />
            <Route exact path="/settings/users" component={SettingsUser} />
          </Switch>
        </div>
        <Footer />
      </>
    );
  }

  return <div className="wrapper">{template}</div>;
};

const mapStateToProps = (state) => ({
  user: state.auth.currentUser,
  token: state.auth.token,
  performance: state.auth.performance,
  connected: getConnected(state.socket),
});

const mapDispatchToProps = (dispatch) => ({
  onCredential: (credentials) =>
    dispatch({ type: ActionTypes.ADD_CREDENTIALS, credentials }),
  onUserLoad: (user) =>
    dispatch({ type: ActionTypes.LOAD_USER, currentUser: user }),
  onUserLogout: () => dispatch({ type: ActionTypes.LOGOUT_USER }),
  onDashboard: (data) =>
    dispatch({ type: ActionTypes.ADD_DATA, performance: data }),

  connect: (url) =>
    dispatch(websocketConnect(url, ActionTypes.WEBSOCKET_PREFIX)),
  disconnect: () => dispatch(disconnect(ActionTypes.WEBSOCKET_PREFIX)),
  onSendMessage: (message) =>
    dispatch(send(message, ActionTypes.WEBSOCKET_PREFIX)),
});

export default connectRedux(mapStateToProps, mapDispatchToProps)(Main);
