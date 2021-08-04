import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import ReactECharts from "echarts-for-react";
import Moment from "react-moment";

import { send } from "@giantmachines/redux-websocket";

import * as ActionTypes from "../store/actions";
import { getConnected } from "../store/reducers/socket";

const Home = ({ connected, performance, onSendMessage, dashboard }) => {
  const [statePerform, updatePerform] = useState({});
  const [lineChart, setLineChart] = useState(null);
  const [taskChart, setTaskChart] = useState(null);

  useEffect(() => {
    let mounted = true;
    if (connected === true) {
      onSendMessage({
        eventItem: "performance",
        eventName: "update",
      });

      if (performance.messages.length > 0 && mounted) {
        const onmessage = performance.messages.filter(
          (msg) => msg.type === "INCOMING"
        );
        if (onmessage[0].data.eventItem === "performance") {
          const eventPayload = onmessage[0].data.eventPayload;
          updatePerform(eventPayload.performance);
          setLineChart(eventPayload.performance.daily);
          setTaskChart(eventPayload.performance.task);
        } else if (onmessage[0].data.eventItem == "open") {
          updatePerform(dashboard);
          setLineChart(dashboard.daily);
          setTaskChart(dashboard.task);
        }
      }
    } else {
      updatePerform(dashboard);
      setLineChart(dashboard.daily);
      setTaskChart(dashboard.task);
    }

    return () => (mounted = false);
  }, [onSendMessage, dashboard]);

  return (
    <>
      <div className="content-header">
        <div className="container-fluid">
          <div className="row mb-2">
            <div className="col-sm-6">
              <h1 className="m-0">Performaces</h1>
            </div>
            <div className="col-sm-6">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Starter Page</li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <section className="content">
        <div className="container-fluid">
          <div className="row">
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box">
                <span className="info-box-icon bg-info elevation-1">
                  <i className="fas fa-flag" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Task TODO</span>
                  <span className="info-box-number">
                    {Object.keys(statePerform).length > 0
                      ? statePerform.card_category[0].count
                      : 0}
                  </span>
                </div>
              </div>
            </div>
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box mb-3">
                <span className="info-box-icon bg-danger elevation-1">
                  <i className="fas fa-thumbs-up" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Task In Progress</span>
                  <span className="info-box-number">
                    {Object.keys(statePerform).length > 0
                      ? statePerform.card_category[1].count
                      : 0}
                  </span>
                </div>
              </div>
            </div>
            <div className="clearfix hidden-md-up" />
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box mb-3">
                <span className="info-box-icon bg-warning elevation-1">
                  <i className="fas fa-vial" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Review/Testing</span>
                  <span className="info-box-number">
                    {Object.keys(statePerform).length > 0
                      ? statePerform.card_category[3].count
                      : 0}
                  </span>
                </div>
              </div>
            </div>
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box mb-3">
                <span className="info-box-icon bg-success elevation-1">
                  <i className="fas fa-check-square" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Task Done</span>
                  <span className="info-box-number">
                    {Object.keys(statePerform).length > 0
                      ? statePerform.card_category[2].count
                      : 0}
                  </span>
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-md-8">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">Distribution</h3>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="remove"
                    >
                      <i className="fas fa-times" />
                    </button>
                  </div>
                </div>
                <div className="card-body p-0">
                  {lineChart ? <ReactECharts option={lineChart} /> : ""}
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">Task</h3>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="remove"
                    >
                      <i className="fas fa-times" />
                    </button>
                  </div>
                </div>
                <div className="card-body" style={{ height: "300px" }}>
                  {taskChart ? <ReactECharts option={taskChart} /> : ""}
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-md-6">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">Recent Activity</h3>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="remove"
                    >
                      <i className="fas fa-times" />
                    </button>
                  </div>
                </div>
                <div
                  className="card-body"
                  style={{ maxHeight: "500px", overflowY: "auto" }}
                >
                  <div className="post">
                    <div className="user-block">
                      <img
                        className="img-circle img-bordered-sm"
                        src="../../dist/img/user1-128x128.jpg"
                        alt="user image"
                      />
                      <span className="username">
                        <a href="#">Jonathan Burke Jr.</a>
                      </span>
                      <span className="description">
                        Shared publicly - 7:45 PM today
                      </span>
                    </div>
                    <p>
                      Lorem ipsum represents a long-held tradition for
                      designers, typographers and the like. Some people hate it
                      and argue for its demise, but others ignore.
                    </p>
                    <p>
                      <a href="#" className="link-black text-sm">
                        <i className="fas fa-link mr-1"></i> Demo File 1 v2
                      </a>
                    </p>
                  </div>
                  <div className="post">
                    <div className="user-block">
                      <img
                        className="img-circle img-bordered-sm"
                        src="../../dist/img/user1-128x128.jpg"
                        alt="user image"
                      />
                      <span className="username">
                        <a href="#">Jonathan Burke Jr.</a>
                      </span>
                      <span className="description">
                        Shared publicly - 7:45 PM today
                      </span>
                    </div>
                    <p>
                      Lorem ipsum represents a long-held tradition for
                      designers, typographers and the like. Some people hate it
                      and argue for its demise, but others ignore.
                    </p>
                    <p>
                      <a href="#" className="link-black text-sm">
                        <i className="fas fa-link mr-1"></i> Demo File 1 v2
                      </a>
                    </p>
                  </div>
                </div>
              </div>
            </div>
            <div className="col-md-6">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">Users Active</h3>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="remove"
                    >
                      <i className="fas fa-times" />
                    </button>
                  </div>
                </div>
                <div className="card-body p-0">
                  <div className="table-responsive">
                    <div id="table-container">
                      <table className="table m-0">
                        <thead>
                          <tr>
                            <th>Name</th>
                            <th>Title</th>
                            <th>Last active</th>
                          </tr>
                        </thead>
                        <tbody>
                          {Object.keys(statePerform).length > 0 ? (
                            statePerform.online_users.map((item, index) => {
                              return (
                                <tr key={index}>
                                  <td>{item.name}</td>
                                  <td>title</td>
                                  <td>
                                    <Moment fromNow>{item.last_active}</Moment>
                                  </td>
                                </tr>
                              );
                            })
                          ) : (
                            <tr />
                          )}
                        </tbody>
                      </table>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
  dashboard: state.dashboard.performance,
  connected: getConnected(state.socket),
  performance: state.socket,
});

const mapDispatchToProps = (dispatch) => ({
  onSendMessage: (message) =>
    dispatch(send(message, ActionTypes.WEBSOCKET_PREFIX)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Home);
