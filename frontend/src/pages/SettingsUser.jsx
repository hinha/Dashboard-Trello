import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import { ToastContainer } from "react-toastify";
import Moment from "react-moment";

import { UPDATE_USER_SETTING } from "../services/setting";
import { FormTrello, FormUser, FormRole } from "../modules/form/SettingsForm";

function Settings({ onClickSidebarApi }) {
  const [getTrelloUser, setTrelloUser] = useState([]);
  const [optTrelloUserSelected, setOptUserTrello] = useState([]);
  const [getAccessControl, setAccessControl] = useState({});
  const [getAccountList, setAccountList] = useState([]);

  useEffect(() => {
    let mounted = true;
    const fetchSetting = async () => {
      const result = await onClickSidebarApi(UPDATE_USER_SETTING);

      let optUser = [];
      if (result.trello) {
        optUser.push({ value: "", label: "Select ID" });
        for (let i = 0; i < result.trello.user.length; i++) {
          const element = result.trello.user[i];
          optUser.push({ value: element.id, label: element.name });
        }
      }
      setOptUserTrello(optUser);
      setAccessControl(result.access);
      setAccountList(result.account);
      setTrelloUser(result.trello.trello_user);
    };
    if (mounted) {
      fetchSetting();
    }

    return () => (mounted = false);
  }, [onClickSidebarApi]);

  const tableListAccount = (newAccount) => {
    if (newAccount.data) {
      setAccountList([...getAccountList, newAccount.data]);
    }
  };

  const tableListUserTrello = (newUser) => {
    if (newUser.data) {
      setTrelloUser([...getTrelloUser, newUser.data]);
    }
  };

  // const trelloAddButton = (e) => {
  //   e.preventDefault();
  //   onSendMessage({
  //     eventItem: "setting:trello",
  //     eventName: "add",
  //   });
  // };

  return (
    <>
      <div className="content-header">
        <div className="container-fluid">
          <div className="row mb-2">
            <div className="col-sm-6">
              <h1 className="m-0">Settings User</h1>
            </div>
            <div className="col-sm-6">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Settings User</li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <section className="content">
        <div className="container-fluid">
          <ToastContainer />
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

                        <FormUser
                          handleRequestAPI={onClickSidebarApi}
                          dispatch={tableListAccount}
                        />
                      </div>
                    </div>
                    <div className="col-md-6">
                      <div className="card card-primary">
                        <div className="card-header">
                          <h3 className="card-title">Assign Role</h3>
                        </div>
                        <FormRole
                          listOptions={getAccessControl}
                          listAccount={getAccountList}
                          handleRequestAPI={onClickSidebarApi}
                        />
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
                            <tbody>
                              {getAccountList &&
                                getAccountList.map((item, index) => {
                                  let status;
                                  if (item.suspend_status === true) {
                                    status = (
                                      <span className="badge bg-danger">
                                        Deactivate
                                      </span>
                                    );
                                  } else {
                                    status = (
                                      <span className="badge bg-primary">
                                        Activate
                                      </span>
                                    );
                                  }
                                  return (
                                    <tr key={index}>
                                      <td>{item.id}</td>
                                      <td>{item.name}</td>
                                      <td>{item.username}</td>
                                      <td>{item.email}</td>
                                      <td>{status}</td>
                                      <td>
                                        <Moment toNow>{item.last_login}</Moment>
                                      </td>
                                      {/* <td>{item.created_at}</td> */}
                                      <td>
                                        <Moment
                                          locale="id"
                                          parse="YYYY-MM-DD HH:mm"
                                          withTitle
                                        >
                                          {item.created_at}
                                        </Moment>
                                      </td>

                                      <td>
                                        <button
                                          type="submit"
                                          className="btn btn-primary btn-sm"
                                          value="${value.id},${value.username}"
                                        >
                                          EDIT
                                        </button>
                                        <button
                                          type="submit"
                                          className="btn btn-danger btn-sm"
                                          value="${value.id},${value.username}"
                                          data-toggle="modal"
                                          data-target="#exampleModal"
                                        >
                                          DELETE
                                        </button>
                                      </td>
                                    </tr>
                                  );
                                })}
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
                                <th>No</th>
                                <th>Board Name</th>
                                <th>Board ID</th>
                                <th>Member ID</th>
                                <th>CreatedAt</th>
                              </tr>
                            </thead>
                            <tbody>
                              {getTrelloUser ? (
                                getTrelloUser.map((item, index) => {
                                  return (
                                    <tr key={index}>
                                      <td>{(index += 1)}</td>
                                      <td>{item.board_name}</td>
                                      <td>{item.board_id}</td>
                                      <td>{item.card_member_id}</td>
                                      <td>
                                        <Moment toNow>{item.created_at}</Moment>
                                      </td>
                                    </tr>
                                  );
                                })
                              ) : (
                                <tr></tr>
                              )}
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
                        <FormTrello
                          dispatch={tableListUserTrello}
                          listOptions={optTrelloUserSelected}
                          handleRequestAPI={onClickSidebarApi}
                        />
                        {/* form end */}
                      </div>
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
}

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(Settings);
