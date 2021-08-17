import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

function SettingsDetail() {
  useEffect(() => {
    console.log("load Settings Detail");
  }, []);

  return (
    <>
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
            <div className="col-md-3">
              {/* Profile Image */}
              <div className="card card-primary card-outline">
                <div className="card-body box-profile">
                  <div className="text-center">
                    <img
                      className="profile-user-img img-fluid img-circle"
                      src="../../dist/img/user4-128x128.jpg"
                      alt="User profile picture"
                    />
                  </div>
                  <h3 className="profile-username text-center">
                    Nina Mcintire
                  </h3>
                  <p className="text-muted text-center">Software Engineer</p>
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
                  <p className="text-muted">
                    B.S. in Computer Science from the University of Tennessee at
                    Knoxville
                  </p>
                  <hr />
                  <strong>
                    <i className="fas fa-map-marker-alt mr-1" /> Location
                  </strong>
                  <p className="text-muted">Malibu, California</p>
                  <hr />
                  <strong>
                    <i className="far fa-file-alt mr-1" /> Notes
                  </strong>
                  <p className="text-muted">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                    Etiam fermentum enim neque.
                  </p>
                </div>
                {/* /.card-body */}
              </div>
              {/* /.card */}
            </div>
            {/* /.col */}
            <div className="col-md-9">
              <div className="card">
                <div className="card-header p-2">
                  <ul className="nav nav-pills">
                    <li className="nav-item">
                      <a
                        className="nav-link active"
                        href="#settings"
                        data-toggle="tab"
                      >
                        Settings
                      </a>
                    </li>
                  </ul>
                </div>
                {/* /.card-header */}
                <div className="card-body">
                  <div className="tab-content">
                    {/* /.tab-pane */}
                    <div className="tab-pane active" id="settings">
                      <form className="form-horizontal">
                        <div className="form-group row">
                          <label
                            htmlFor="inputName"
                            className="col-sm-2 col-form-label"
                          >
                            Name
                          </label>
                          <div className="col-sm-10">
                            <input
                              type="email"
                              className="form-control"
                              id="inputName"
                              placeholder="Name"
                            />
                          </div>
                        </div>
                        <div className="form-group row">
                          <label
                            htmlFor="inputEmail"
                            className="col-sm-2 col-form-label"
                          >
                            Email
                          </label>
                          <div className="col-sm-10">
                            <input
                              type="email"
                              className="form-control"
                              id="inputEmail"
                              placeholder="Email"
                            />
                          </div>
                        </div>
                        <div className="form-group row">
                          <label
                            htmlFor="inputName2"
                            className="col-sm-2 col-form-label"
                          >
                            Name
                          </label>
                          <div className="col-sm-10">
                            <input
                              type="text"
                              className="form-control"
                              id="inputName2"
                              placeholder="Name"
                            />
                          </div>
                        </div>
                        <div className="form-group row">
                          <label
                            htmlFor="inputExperience"
                            className="col-sm-2 col-form-label"
                          >
                            Experience
                          </label>
                          <div className="col-sm-10">
                            <textarea
                              className="form-control"
                              id="inputExperience"
                              placeholder="Experience"
                              defaultValue={""}
                            />
                          </div>
                        </div>
                        <div className="form-group row">
                          <label
                            htmlFor="inputSkills"
                            className="col-sm-2 col-form-label"
                          >
                            Skills
                          </label>
                          <div className="col-sm-10">
                            <input
                              type="text"
                              className="form-control"
                              id="inputSkills"
                              placeholder="Skills"
                            />
                          </div>
                        </div>
                        <div className="form-group row">
                          <div className="offset-sm-2 col-sm-10">
                            <div className="checkbox">
                              <label>
                                <input type="checkbox" /> I agree to the{" "}
                                <a href="#">terms and conditions</a>
                              </label>
                            </div>
                          </div>
                        </div>
                        <div className="form-group row">
                          <div className="offset-sm-2 col-sm-10">
                            <button type="submit" className="btn btn-danger">
                              Submit
                            </button>
                          </div>
                        </div>
                      </form>
                    </div>
                    {/* /.tab-pane */}
                  </div>
                  {/* /.tab-content */}
                </div>
                {/* /.card-body */}
              </div>
              {/* /.card */}
            </div>
            {/* /.col */}
          </div>
        </div>
      </section>
    </>
  );
}

const mapStateToProps = (state) => ({
  socket: state.socket,
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(SettingsDetail);
