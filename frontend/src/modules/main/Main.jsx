import React, { useState, useEffect } from "react";
import { Route, Switch } from "react-router-dom";
import { connect } from "react-redux";
// import useWebSocket, { ReadyState } from "react-use-websocket";

import Footer from "./footer/Footer";
import Header from "./header/Header";
import Sidebar from "./../../components/layout/Sidebar";
import PageLoading from "./../../components/page-loading/PageLoading";
import Attendence from "./../../pages/Attendence";
import Home from "./../../pages/Home";

import * as AuthService from "./../../services/profile";
import * as ActionTypes from "../../store/actions";
import Socket from "./socket";

const Main = ({ token, onSocket, onUserLoad, onUserLogout }) => {
  const [appLoadingState, updateAppLoading] = useState(false);
  const [menusidebarState, updateMenusidebarState] = useState({
    isMenuSidebarCollapsed: false,
  });

  const [getToken, setToken] = useState("");

  useEffect(() => {
    updateAppLoading(true);
    const fetchProfile = async () => {
      try {
        const response = await AuthService.getProfile(token);

        onUserLoad({ ...response });
        setToken(response.credentials);
        updateAppLoading(false);
      } catch (error) {
        console.log(error);
        if (error.response) {
          if (error.response.status === 401) {
            onUserLogout();
          }
        }
        updateAppLoading(true);
      }
    };

    fetchProfile();

    return () => {};
  }, [onUserLoad, onSocket]);

  useEffect(() => {
    try {
      if (getToken !== "") {
        const Sock = new Socket("userID", getToken);
        Sock.subscribeToSocketMessage();
        onSocket(Sock);
      }
    } catch (e) {
      console.error(e, "errr");
    }
    return () => {};
  }, [getToken]);

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
});

const mapDispatchToProps = (dispatch) => ({
  onUserLoad: (user) =>
    dispatch({ type: ActionTypes.LOAD_USER, currentUser: user }),
  onUserLogout: () => dispatch({ type: ActionTypes.LOGOUT_USER }),
  onSocket: (socket) =>
    dispatch({ type: ActionTypes.ADD_SOCKET, socket: socket }),
});

export default connect(mapStateToProps, mapDispatchToProps)(Main);
