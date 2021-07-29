// import React, { useState } from "react";
import React from "react";
import { Link, NavLink } from "react-router-dom";
import { connect } from "react-redux";

import {
  SidebarDataPermormance,
  SidebarDataAnalytics,
  SidebarDataSettings,
} from "../SidebarData";

function Sidebar({ user }) {
  return (
    <aside className="main-sidebar sidebar-dark-primary elevation-4">
      {/* Brand Logo */}
      <Link to="/" className="brand-link">
        <img
          src="../dist/img/AdminLTELogo.png"
          alt="AdminLTE Logo"
          className="brand-image img-circle elevation-3"
          style={{ opacity: ".8" }}
        />
        <span className="brand-text font-weight-light">AdminLTE 3</span>
      </Link>
      {/* Sidebar */}
      <div className="sidebar">
        {/* Sidebar user panel (optional) */}
        <div className="user-panel mt-3 pb-3 mb-3 d-flex">
          <div className="image">
            <img
              src="../dist/img/user2-160x160.jpg"
              className="img-circle elevation-2"
              alt="User Image"
            />
          </div>
          <div className="info">
            <a href="#" className="d-block">
              {user.name}
            </a>
          </div>
        </div>
        {/* Sidebar Menu */}
        <nav className="mt-2">
          <ul
            className="nav nav-pills nav-sidebar flex-column"
            data-widget="treeview"
            role="menu"
            data-accordion="false"
          >
            <li className="nav-item has-treeview menu-open">
              <NavLink
                className="nav-link"
                activeClassName="active"
                to="/dashboard"
              >
                <i className="nav-icon fas fa-tachometer-alt" />
                <p>
                  Dashboard <i className="fas fa-angle-left right"></i>
                </p>
              </NavLink>
              <ul className="nav nav-treeview">
                {SidebarDataPermormance.map((item, index) => {
                  return (
                    <li className="nav-item" key={index}>
                      <NavLink
                        exact
                        to={item.path}
                        activeClassName="active"
                        className="nav-link"
                      >
                        {item.icon}
                        <p>{item.title}</p>
                      </NavLink>
                    </li>
                  );
                })}
              </ul>
            </li>
            <li className="nav-item has-treeview menu-open">
              <NavLink
                className="nav-link"
                activeClassName="active"
                to="/analytics"
              >
                <i className="nav-icon fas fa-chart-pie"></i>
                <p>
                  Analytics<i className="fas fa-angle-left right"></i>
                </p>
              </NavLink>
              <ul className="nav nav-treeview">
                {SidebarDataAnalytics.map((item, index) => {
                  return (
                    <li className="nav-item" key={index}>
                      <NavLink
                        exact
                        to={item.path}
                        activeClassName="active"
                        className="nav-link"
                      >
                        {item.icon}
                        <p>{item.title}</p>
                      </NavLink>
                    </li>
                  );
                })}
              </ul>
            </li>
            <li className="nav-item has-treeview menu-open">
              <NavLink
                className="nav-link"
                activeClassName="active"
                to="/settings"
              >
                <i className="nav-icon fas fa-cog"></i>
                <p>
                  Settings
                  <i className="fas fa-angle-left right"></i>
                </p>
              </NavLink>
              <ul className="nav nav-treeview">
                {SidebarDataSettings.map((item, index) => {
                  return (
                    <li className="nav-item" key={index}>
                      <NavLink
                        exact
                        to={item.path}
                        activeClassName="active"
                        className="nav-link"
                      >
                        {item.icon}
                        <p>{item.title}</p>
                      </NavLink>
                    </li>
                  );
                })}
              </ul>
            </li>
          </ul>
        </nav>
        {/* /.sidebar-menu */}
      </div>
      {/* /.sidebar */}
    </aside>
  );
}

const mapStateToProps = (state) => ({
  user: state.auth.currentUser,
});

export default connect(mapStateToProps, null)(Sidebar);
