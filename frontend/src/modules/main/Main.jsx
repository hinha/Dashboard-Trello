import React, { useState, useEffect } from "react";
import { Route, Switch } from "react-router-dom";
import { connect as connectRedux } from "react-redux";

import Footer from "./footer/Footer";
import Header from "./header/Header";
import Sidebar from "./../../components/layout/Sidebar";
import PageLoading from "./../../components/page-loading/PageLoading";
import Home from "./../../pages/Home";
import Attendence from "./../../pages/Attendence";
import KMedoids from "../../pages/KMedoids";
import TrelloBoard from "../../pages/TrelloBoard";
import SettingsDetail from "../../pages/SettingsDetail";
import SettingsUser from "../../pages/SettingsUser";

import * as AuthService from "./../../services/profile";
import * as DashboardService from "./../../services/dashboard";
import * as AnalyticService from "./../../services/analytics";
import * as SettingService from "./../../services/setting";
import * as ActionTypes from "../../store/actions";

const Main = ({ token, onCredential, onUserLoad, onArn, onUserLogout }) => {
  const [appLoadingState, updateAppLoading] = useState(false);
  const [menusidebarState, updateMenusidebarState] = useState({
    isMenuSidebarCollapsed: false,
  });
  const [arnState, updateArnState] = useState([]);
  const [getCredential, setCredential] = useState("");

  // const [getToken, setToken] = useState("");

  useEffect(() => {
    updateAppLoading(true);
    let mounted = true;
    const fetchProfile = async () => {
      try {
        const response = await AuthService.getProfile(token);

        onUserLoad({ ...response });
        onCredential(response.credentials);
        setCredential(response.credentials);

        let newArn = [];
        for (var i = 0; i < response.arn.length; i++) {
          newArn.push(response.arn[i].toLowerCase());
        }
        onArn(newArn);

        if (mounted) {
          updateAppLoading(false);
          updateArnState(response.arn);
        }
      } catch (error) {
        // console.log(error);
        if (error.response) {
          if (error.response.status === 401) {
            onUserLogout();
          }
        }
      }
    };

    fetchProfile();

    return () => (mounted = false);
  }, [onUserLoad, onArn, onCredential]);

  const toggleMenuSidebar = () => {
    updateMenusidebarState({
      isMenuSidebarCollapsed: !menusidebarState.isMenuSidebarCollapsed,
    });
  };

  const onClickSidebarApi = async (item, body = {}, params = "") => {
    let data;
    if (item == "performance") {
      data = await DashboardService.getDashboard(token, getCredential);
    } else if (item == AnalyticService.UPDATE_ANALYTIC_TRELLO) {
      data = await AnalyticService.getTrelloCard(token, getCredential);
    } else if (item == SettingService.UPDATE_USER_SETTING) {
      data = await SettingService.getSettingUser(token, getCredential);
    } else if (item == SettingService.ADD_USER_SETTING) {
      data = await SettingService.addSettingUser(token, getCredential, body);
    } else if (item == SettingService.EDIT_USER_SETTING) {
      data = await SettingService.editSettingUser(token, getCredential, body);
    } else if (item == SettingService.DEL_USER_SETTING) {
      data = await SettingService.delSettingUser(token, getCredential, params);
    } else if (item == SettingService.ROLE_USER_SETTING) {
      data = await SettingService.roleSettingUser(token, getCredential, body);
    } else if (item == SettingService.TRELLO_USER_SETTING) {
      data = await SettingService.trelloSettingUser(token, getCredential, body);
    }

    return data;
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

  let Router = (
    <Switch>
      {arnState ? (
        arnState.map((item, key) => {
          item = item.toLowerCase();
          const menu = item.split(":");

          let route;
          switch (menu[0]) {
            case "dashboard":
              if (menu[1] === "performance") {
                route = (
                  <Route
                    exact
                    path="/dashboard"
                    render={(props) => (
                      <Home {...props} onClickSidebarApi={onClickSidebarApi} />
                    )}
                    key={key}
                  />
                );
              } else if (menu[1] === "attendance") {
                route = (
                  <Route
                    exact
                    path="/dashboard/attendence"
                    render={(props) => (
                      <Attendence
                        {...props}
                        onClickSidebarApi={onClickSidebarApi}
                      />
                    )}
                    key={key}
                  />
                );
              } else {
                <Route
                  exact
                  path="/dashboard/employee"
                  component={Attendence}
                  key={key}
                />;
              }
              break;
            case "user":
              if (menu[1] === "detail") {
                route = (
                  <Route
                    exact
                    path="/settings"
                    component={SettingsDetail}
                    key={key}
                  />
                );
              } else if (menu[1] === "manage") {
                route = (
                  <Route
                    exact
                    path="/settings/users"
                    render={(props) => (
                      <SettingsUser
                        {...props}
                        onClickSidebarApi={onClickSidebarApi}
                      />
                    )}
                    key={key}
                  />
                );
              }
              break;
            case "analytics":
              route = (
                <Route exact path="/analytics" component={KMedoids} key={key} />
              );
              break;
            case "trello":
              route = (
                <Route
                  exact
                  path="/analytics/trello"
                  render={(props) => (
                    <TrelloBoard
                      {...props}
                      onClickSidebarApi={onClickSidebarApi}
                    />
                  )}
                  key={key}
                />
              );
              break;
          }

          if (route) {
            return route;
          }
        })
      ) : (
        <div />
      )}
    </Switch>
  );
  let template;

  if (appLoadingState) {
    template = <PageLoading />;
  } else {
    template = (
      <>
        <Header toggleMenuSidebar={toggleMenuSidebar} />
        <Sidebar />
        {Router}
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
  onCredential: (credentials) =>
    dispatch({ type: ActionTypes.ADD_CREDENTIALS, credentials }),
  onUserLoad: (user) =>
    dispatch({ type: ActionTypes.LOAD_USER, currentUser: user }),
  onUserLogout: () => dispatch({ type: ActionTypes.LOGOUT_USER }),

  onArn: (arnList) => dispatch({ type: ActionTypes.ARN_USER, arnList }),
});

export default connectRedux(mapStateToProps, mapDispatchToProps)(Main);
