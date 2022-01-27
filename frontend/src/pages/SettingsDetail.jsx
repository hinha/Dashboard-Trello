import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

import { DETAIL_USER_SETTING } from "../services/setting";

function SettingsDetail({ user, onClickSidebarApi }) {
  const [getProfilName, setProfilName] = useState("");
  const [getJobTitle, setJobTitle] = useState("");

  useEffect(() => {
    let mounted = true;
    const fetchSetting = async () => {
      const result = await onClickSidebarApi(DETAIL_USER_SETTING);
      console.log(result);
      setJobTitle(result.job_title);
    };

    if (mounted) {
      setProfilName(user.name);
      fetchSetting();
    }

    return () => (mounted = false);
  }, []);

  return (
    <div className="content-wrapper">
      <div className="content-header">
        <div className="container-fluid">
          <div className="row mb-2">
            <div className="col-sm-6">
              <h1 className="m-0">Settings Detail</h1>
            </div>
            <div className="col-sm-6">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Settings Detail</li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <section className="content">
        <div className="container-fluid">
          <div className="row">
            <div className="col-md-12">
              {/* Profile Image */}
              <div className="card card-primary card-outline">
                <div className="card-body box-profile">
                  <div className="text-center">
                    <img
                      className="profile-user-img img-fluid img-circle"
                      src="../../dist/img/user2-160x160.jpg"
                      alt="User profile picture"
                    />
                  </div>
                  <h3 className="profile-username text-center">
                    {getProfilName}
                  </h3>
                  <p className="text-muted text-center">{getJobTitle}</p>
                </div>
                {/* /.card-body */}
              </div>
              {/* /.card */}
              {/* About Me Box */}
              <div className="card card-primary">
                <div className="card-header">
                  <h3 className="card-title">About Me</h3>
                </div>
                {/* /.card-header */}
                <div className="card-body">
                  <strong>
                    <i className="fas fa-book mr-1" /> Education
                  </strong>
                  <p className="text-muted">-</p>
                  <hr />
                  <strong>
                    <i className="fas fa-map-marker-alt mr-1" /> Location
                  </strong>
                  <p className="text-muted">-</p>
                  <hr />
                  <strong>
                    <i className="far fa-file-alt mr-1" /> Notes
                  </strong>
                  <p className="text-muted">-</p>
                </div>
                {/* /.card-body */}
              </div>
              {/* /.card */}
            </div>
            {/* /.col */}
          </div>
        </div>
      </section>
    </div>
  );
}

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
  user: state.auth.currentUser,
});

export default connect(mapStateToProps, null)(SettingsDetail);
