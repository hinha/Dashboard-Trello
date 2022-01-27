import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import ReactECharts from "echarts-for-react";
import Moment from "react-moment";

const Home = ({ onClickSidebarApi }) => {
  const [statePerform, updatePerform] = useState({});
  const [lineChart, setLineChart] = useState(null);
  const [taskChart, setTaskChart] = useState(null);

  useEffect(() => {
    let mounted = true;

    const fetchPerformance = async () => {
      const result = await onClickSidebarApi("performance");
      updatePerform(result);
      setLineChart(result.daily);
      setTaskChart(result.task);
    };

    if (mounted) {
      fetchPerformance();
    }

    return () => (mounted = false);
  }, [onClickSidebarApi]);

  return (
    <div className="content-wrapper">
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
                <li className="breadcrumb-item active">Dashboard</li>
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
                  {statePerform.card_activity &&
                    statePerform.card_activity.map((item, index) => {
                      return (
                        <>
                          <div className="post" key={index}>
                            <div className="user-block">
                              <span className="username">
                                {item.card_member_name}
                              </span>
                              <span className="description">
                                {item.card_created_at}
                              </span>
                            </div>
                            <p>{item.card_name}</p>
                            <p>
                              <a
                                className="link-black text-sm"
                                onClick={() => window.open(item.card_url)}
                              >
                                <i className="fas fa-link mr-1"></i>[
                                {item.card_category}] See Detail
                              </a>
                            </p>
                          </div>
                        </>
                      );
                    })}
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
                            <th>Last active</th>
                          </tr>
                        </thead>
                        <tbody>
                          {statePerform &&
                            statePerform.online_users != null &&
                            statePerform.online_users.map((item, index) => {
                              return (
                                <tr key={index}>
                                  <td>{item.name}</td>
                                  <td>
                                    <Moment fromNow>{item.last_active}</Moment>
                                  </td>
                                </tr>
                              );
                            })}
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
    </div>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(Home);
