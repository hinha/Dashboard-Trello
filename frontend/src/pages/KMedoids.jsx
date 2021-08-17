import React from "react";
import { Link } from "react-router-dom";

const KMedoids = () => {
  return (
    <>
      <div className="content-header">
        <div className="container-fluid">
          <div className="row mb-2">
            <div className="col-sm-6">
              <h1 className="m-0">K-Medoids</h1>
            </div>
            <div className="col-sm-6">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Analytics</li>
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
                  <span className="info-box-text">Total TODO</span>
                  <span className="info-box-number">0</span>
                </div>
              </div>
            </div>
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box mb-3">
                <span className="info-box-icon bg-danger elevation-1">
                  <i className="fas fa-thumbs-up" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Total In Progress</span>
                  <span className="info-box-number">0</span>
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
                  <span className="info-box-text">Total Review/Testing</span>
                  <span className="info-box-number">0</span>
                </div>
              </div>
            </div>
            <div className="col-12 col-sm-6 col-md-3">
              <div className="info-box mb-3">
                <span className="info-box-icon bg-success elevation-1">
                  <i className="fas fa-check-square" />
                </span>
                <div className="info-box-content">
                  <span className="info-box-text">Total Done</span>
                  <span className="info-box-number">0</span>
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-md-12">
              <div className="card">
                <div className="card-header">
                  <h5 className="card-title">Total Data</h5>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                  </div>
                </div>
                {/* /.card-header */}
                <div className="card-body">
                  <div className="row">
                    <div className="col-md-5">
                      <table className="table table-bordered">
                        <thead>
                          <tr>
                            <th style={{ width: 10 }}>#</th>
                            <th>Task</th>
                            <th>Progress</th>
                            <th style={{ width: 40 }}>Label</th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr>
                            <td>1.</td>
                            <td>Update software</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar progress-bar-danger"
                                  style={{ width: "55%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-danger">55%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>2.</td>
                            <td>Clean database</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar bg-warning"
                                  style={{ width: "70%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-warning">70%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>2.</td>
                            <td>Clean database</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar bg-warning"
                                  style={{ width: "70%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-warning">70%</span>
                            </td>
                          </tr>

                          <tr>
                            <td>3.</td>
                            <td>Cron job running</td>
                            <td>
                              <div className="progress progress-xs progress-striped active">
                                <div
                                  className="progress-bar bg-primary"
                                  style={{ width: "30%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-primary">30%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>4.</td>
                            <td>Fix and squish bugs</td>
                            <td>
                              <div className="progress progress-xs progress-striped active">
                                <div
                                  className="progress-bar bg-success"
                                  style={{ width: "90%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-success">90%</span>
                            </td>
                          </tr>
                        </tbody>
                      </table>
                      <ul className="pagination pagination-sm m-0 float-right">
                        <li className="page-item">
                          <a className="page-link" href="#">
                            «
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            1
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            2
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            3
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            »
                          </a>
                        </li>
                      </ul>
                    </div>
                    {/* /.col */}
                    <div className="col-md-7">
                      <div className="chart">
                        <div className="chartjs-size-monitor">
                          <div className="chartjs-size-monitor-expand">
                            <div className />
                          </div>
                          <div className="chartjs-size-monitor-shrink">
                            <div className />
                          </div>
                        </div>
                        <p>Line Chart</p>
                      </div>
                    </div>
                    {/* /.col */}
                  </div>
                  {/* /.row */}
                </div>
                {/* ./card-body */}
              </div>
              {/* /.card */}
            </div>
            {/* /.col */}
          </div>
          <div className="row">
            <div className="col-md-12">
              <div className="card">
                <div className="card-header">
                  <h5 className="card-title">Cluster</h5>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                  </div>
                </div>
                {/* /.card-header */}
                <div className="card-body">
                  <div className="row">
                    <div className="col-md-5">
                      <table className="table table-bordered">
                        <thead>
                          <tr>
                            <th style={{ width: 10 }}>#</th>
                            <th>Task</th>
                            <th>Progress</th>
                            <th style={{ width: 40 }}>Label</th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr>
                            <td>1.</td>
                            <td>Update software</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar progress-bar-danger"
                                  style={{ width: "55%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-danger">55%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>2.</td>
                            <td>Clean database</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar bg-warning"
                                  style={{ width: "70%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-warning">70%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>2.</td>
                            <td>Clean database</td>
                            <td>
                              <div className="progress progress-xs">
                                <div
                                  className="progress-bar bg-warning"
                                  style={{ width: "70%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-warning">70%</span>
                            </td>
                          </tr>

                          <tr>
                            <td>3.</td>
                            <td>Cron job running</td>
                            <td>
                              <div className="progress progress-xs progress-striped active">
                                <div
                                  className="progress-bar bg-primary"
                                  style={{ width: "30%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-primary">30%</span>
                            </td>
                          </tr>
                          <tr>
                            <td>4.</td>
                            <td>Fix and squish bugs</td>
                            <td>
                              <div className="progress progress-xs progress-striped active">
                                <div
                                  className="progress-bar bg-success"
                                  style={{ width: "90%" }}
                                />
                              </div>
                            </td>
                            <td>
                              <span className="badge bg-success">90%</span>
                            </td>
                          </tr>
                        </tbody>
                      </table>
                      <ul className="pagination pagination-sm m-0 float-right">
                        <li className="page-item">
                          <a className="page-link" href="#">
                            «
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            1
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            2
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            3
                          </a>
                        </li>
                        <li className="page-item">
                          <a className="page-link" href="#">
                            »
                          </a>
                        </li>
                      </ul>
                    </div>
                    {/* /.col */}
                    <div className="col-md-7">
                      <div className="chart">
                        <div className="chartjs-size-monitor">
                          <div className="chartjs-size-monitor-expand">
                            <div className />
                          </div>
                          <div className="chartjs-size-monitor-shrink">
                            <div className />
                          </div>
                        </div>
                        <p>Line Chart</p>
                      </div>
                    </div>
                    {/* /.col */}
                  </div>
                  {/* /.row */}
                </div>
                {/* ./card-body */}
              </div>
              {/* /.card */}
            </div>
            {/* /.col */}
          </div>

          <div className="row">
            <div className="col-md-12">
              <div className="card">
                <div className="card-header">
                  <h5 className="card-title">Silhouette Coefficient</h5>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                  </div>
                </div>
                {/* /.card-header */}
                <div className="card-body">Bar Chart</div>
              </div>
            </div>
          </div>

          <div className="row">
            <div className="col-md-12">
              <div className="card">
                <div className="card-header">
                  <h5 className="card-title">Hasil Analisis</h5>
                  <div className="card-tools">
                    <button
                      type="button"
                      className="btn btn-tool"
                      data-card-widget="collapse"
                    >
                      <i className="fas fa-minus" />
                    </button>
                  </div>
                </div>
                {/* /.card-header */}
                <div className="card-body">Bar Chart</div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default KMedoids;
