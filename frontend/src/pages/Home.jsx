import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

const Home = ({ socket, dashboard }) => {
  const [statePerform, updatePerform] = useState([]);

  useEffect(() => {
    let mounted = true;
    if (socket !== null) {
      if (socket.socket !== null) {
        const item = socket.socket.tesKiremClick("performance");
        if (mounted) {
          updatePerform(item);
        }
      }
    }
    let cardData = [];
    if (statePerform.length === 0) {
      if (Object.keys(dashboard).length === 0) {
        cardData = [];
      } else {
        cardData = dashboard.card_category;
        console.log(cardData);
      }
    } else {
      cardData = statePerform;
    }

    updatePerform(cardData);
    // return () => (mounted = false);
  }, [socket, dashboard]);

  console.log(statePerform);

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
                    {statePerform.length > 0 ? statePerform[0].count : 0}
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
                    {statePerform.length > 0 ? statePerform[1].count : 0}
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
                    {statePerform.length > 0 ? statePerform[3].count : 0}
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
                    {statePerform.length > 0 ? statePerform[2].count : 0}
                  </span>
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-md-8">
              <div className="card">
                <div className="card-body">
                  <h4 className="mt-0 header-title mb-4">Daily Task</h4>
                  <div id="distSentiment" />
                </div>
              </div>
            </div>
            <div className="col-md-4">
              <div className="card">
                <div className="card-body">
                  <h4 className="mt-0 header-title mb-4">Task</h4>
                  <div id="pieSentiment" />
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-md-6">
              <div className="card">
                <div className="card-body">
                  <h4 className="mt-0 header-title mb-4">Gant cahrt</h4>
                  <div id="gantCahrt" />
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
                            <th>Status</th>
                            <th>Title</th>
                          </tr>
                        </thead>
                        <tbody />
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
  socket: state.socket,
  credentials: state.auth.credentials,
  token: state.auth.token,
  dashboard: state.dashboard.performance,
});

export default connect(mapStateToProps, null)(Home);
