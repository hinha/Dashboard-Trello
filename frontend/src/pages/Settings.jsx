import React from "react";
import { Link } from "react-router-dom";

function Settings() {
  return (
    <>
      <div className="content-header">
        <div className="container-fluid">
          <div className="row mb-2">
            <div className="col-sm-6">
              <h1 className="m-0">Settings</h1>
            </div>
            <div className="col-sm-6">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Settings</li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <section className="content">
        <div className="container-fluid">
          <div className="row">
            <div className="col-md-12">
              <div className="row">
                <div className="col-xs-12 col-md-12 col-lg-12">
                  <ul className="nav nav-pills">
                    <li className="nav-item">
                      <a
                        href="#user"
                        className="nav-link active"
                        data-toggle="tab"
                      >
                        User
                      </a>
                    </li>
                    <li className="nav-item">
                      <a href="#trello" className="nav-link" data-toggle="tab">
                        Trello
                      </a>
                    </li>
                    <li className="nav-item">
                      <a href="#log" className="nav-link" data-toggle="tab">
                        Log
                      </a>
                    </li>
                  </ul>
                </div>
              </div>
              <div className="tab-content">
                <div className="tab-pane fade show active" id="user">
                  <div className="row">
                    <div className="col-md-6">
                      <div className="card card-primary">
                        <div className="card-header">
                          <h3 className="card-title">Register User</h3>
                        </div>

                        <form id="register-account">
                          <input
                            type="hidden"
                            name="csrf"
                            defaultValue="{{ .Token }}"
                          />
                          <div className="card-body">
                            <div className="form-group">
                              <label>Full Name</label>
                              <input
                                type="text"
                                className="form-control"
                                name="name"
                                placeholder="Enter Full Name"
                              />
                            </div>
                            <div className="form-group">
                              <label>Username</label>
                              <input
                                type="text"
                                className="form-control"
                                name="username"
                                placeholder="Enter Username"
                              />
                            </div>
                            <div className="form-group">
                              <label>Email address</label>
                              <input
                                type="email"
                                className="form-control"
                                name="email"
                                placeholder="Enter email"
                              />
                            </div>
                            <div className="form-group">
                              <label>Password</label>
                              <input
                                type="password"
                                className="form-control"
                                name="password"
                                placeholder="Password"
                              />
                            </div>
                          </div>
                          {/* /.card-body */}
                          <div className="card-footer">
                            <button type="submit" className="btn btn-primary">
                              Submit
                            </button>
                          </div>
                        </form>
                      </div>
                    </div>
                    <div className="col-md-6">
                      <div className="card card-primary">
                        <div className="card-header">
                          <h3 className="card-title">Assign Role</h3>
                        </div>

                        {/* form start */}
                        <form id="assign-role">
                          <div className="card-body">
                            <div className="form-group">
                              <label>Select User</label>
                              <select
                                className="form-control select2"
                                style={{ width: "100%" }}
                                id="opt-select-user"
                              ></select>
                            </div>
                            <div className="form-group">
                              <label>Select Role</label>
                              <select
                                className="form-control select2"
                                style={{ width: "100%" }}
                                id="opt-select-role"
                              ></select>
                            </div>
                            <div className="form-group">
                              <label>Select Permission</label>
                              <select
                                className="form-control select2"
                                style={{ width: "100%" }}
                                id="opt-select-permission"
                              ></select>
                            </div>
                          </div>
                          {/* /.card-body */}
                          <div className="card-footer">
                            <button type="submit" className="btn btn-primary">
                              Create
                            </button>
                          </div>
                        </form>
                      </div>
                    </div>
                  </div>

                  <div className="row">
                    <div className="col-md-12">
                      <div className="card">
                        <div className="card-header">
                          <h3 className="card-title">Lists user</h3>

                          <div className="card-tools">
                            <div
                              className="input-group input-group-sm"
                              style={{ width: "150px" }}
                            >
                              <div
                                className="input-group-append"
                                style={{
                                  marginLeft: "-72px",
                                  marginRight: "14px",
                                }}
                              >
                                <button
                                  type="submit"
                                  className="btn btn-primary"
                                >
                                  <i className="fas fa-sync-alt"></i>
                                </button>
                              </div>
                              <input
                                type="text"
                                name="table_search"
                                className="form-control float-right"
                                placeholder="Search"
                              />
                              <div className="input-group-append">
                                <button
                                  type="submit"
                                  className="btn btn-default"
                                >
                                  <i className="fas fa-search"></i>
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>
                        {/* <!-- /.card-header --> */}
                        <div
                          className="card-body table-responsive p-0"
                          style={{ height: "362px" }}
                          id="user-list"
                        >
                          <table className="table table-head-fixed text-nowrap">
                            <thead>
                              <tr>
                                <th>ID</th>
                                <th>Name</th>
                                <th>Username</th>
                                <th>Email</th>
                                <th>Suspend</th>
                                <th>Last Login</th>
                                <th>Created At</th>
                                <th>Action</th>
                              </tr>
                            </thead>
                            <tbody id="user">
                              <tr>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td>
                                  <div className="text-center">
                                    <div
                                      className="spinner-border"
                                      role="status"
                                    >
                                      <span className="sr-only">
                                        Loading...
                                      </span>
                                    </div>
                                  </div>
                                </td>
                                <td></td>
                                <td></td>
                              </tr>
                            </tbody>
                          </table>
                        </div>
                        {/* <!-- /.card-body --> */}
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="trello">
                  <div className="row">
                    <div className="col-md-8">
                      <div className="card">
                        <div className="card-header">
                          <h3 className="card-title">Lists user</h3>

                          <div className="card-tools">
                            <div
                              className="input-group input-group-sm"
                              style={{ width: "150px" }}
                            >
                              <input
                                type="text"
                                name="table_search"
                                className="form-control float-right"
                                placeholder="Search"
                              />

                              <div className="input-group-append">
                                <button
                                  type="submit"
                                  className="btn btn-default"
                                >
                                  <i className="fas fa-search"></i>
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>
                        {/* <!-- /.card-header --> */}
                        <div
                          className="card-body table-responsive p-0"
                          style={{ height: "362px" }}
                        >
                          <table className="table table-head-fixed text-nowrap">
                            <thead>
                              <tr>
                                <th>ID</th>
                                <th>User</th>
                                <th>Date</th>
                                <th>Status</th>
                                <th>Reason</th>
                              </tr>
                            </thead>
                            <tbody>
                              <tr>
                                <td>183</td>
                                <td>John Doe</td>
                                <td>11-7-2014</td>
                                <td>
                                  <span className="tag tag-success">
                                    Approved
                                  </span>
                                </td>
                                <td>
                                  Bacon ipsum dolor sit amet salami venison
                                  chicken flank fatback doner.
                                </td>
                              </tr>
                            </tbody>
                          </table>
                        </div>
                        {/* <!-- /.card-body --> */}
                      </div>
                      {/* <!-- /.card --> */}
                    </div>
                    <div className="col-md-4">
                      <div className="card card-primary">
                        <div className="card-header">
                          <h3 className="card-title">ADD</h3>
                        </div>

                        {/* form start */}
                        <form id="assign-role">
                          <div className="card-body">
                            <div className="form-group">
                              <label>Select User</label>
                              <select
                                className="form-control select2"
                                style={{ width: "100%" }}
                                id="opt-select-user"
                              ></select>
                            </div>
                            <div className="form-group">
                              <label>Board Name</label>
                              <input
                                type="text"
                                className="form-control"
                                name="board_name"
                                placeholder="Enter Name"
                              />
                            </div>
                            <div className="form-group">
                              <label>Member ID</label>
                              <input
                                type="text"
                                className="form-control"
                                name="member_id"
                                placeholder="Enter Id"
                              />
                            </div>
                          </div>
                          {/* /.card-body */}
                          <div className="card-footer">
                            <button type="submit" className="btn btn-primary">
                              Create
                            </button>
                          </div>
                        </form>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="log">
                  <p>Messages tab content ...</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}

export default Settings;
