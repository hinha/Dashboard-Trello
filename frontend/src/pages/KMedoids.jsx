import React, { useEffect } from "react";
import { Link } from "react-router-dom";

import { connect } from "react-redux";

import Clusters from "../components/chart/Clusters";
import ProgressData from "../components/chart/ProgressData";

const KMedoids = () => {
  useEffect(() => {
    let mounted = true;

    if (mounted) {
      console.log("mounted");
    }

    return () => (mounted = false);
  }, []);

  return (
    <div className="content-wrapper">
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
          <ProgressData />
          <Clusters />

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
    </div>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(KMedoids);
