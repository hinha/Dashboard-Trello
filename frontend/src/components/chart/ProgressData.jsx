import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import Spinner from "react-bootstrap/Spinner";
import ReactECharts from "echarts-for-react";

import { getOverviewData } from "../../services/analytics";

const ProgressData = ({ credentials, token }) => {
  const [getDailyCountTask, setDailyCountTask] = useState({});
  const [getCardCategory, setCardCategory] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    let mounted = true;

    const fetchAPI = async () => {
      setLoading(true);
      const result = await getOverviewData(token, credentials, "2021");
      setDailyCountTask(result.daily);
      setCardCategory(result.card_category);
      setLoading(false);
    };

    if (mounted) {
      fetchAPI();
    }

    return () => (mounted = false);
  }, [credentials, token]);

  return (
    <>
      <div className="row">
        <div className="col-12 col-sm-6 col-md-3">
          <div className="info-box">
            <span className="info-box-icon bg-info elevation-1">
              <i className="fas fa-flag" />
            </span>
            <div className="info-box-content">
              <span className="info-box-text">Total TODO</span>
              <span className="info-box-number">
                {loading ? (
                  <Spinner animation="border" size="sm" />
                ) : getCardCategory.length > 0 ? (
                  getCardCategory[0].count
                ) : (
                  0
                )}
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
              <span className="info-box-text">Total In Progress</span>
              <span className="info-box-number">
                {loading ? (
                  <Spinner animation="border" size="sm" />
                ) : getCardCategory.length > 0 ? (
                  getCardCategory[1].count
                ) : (
                  0
                )}
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
              <span className="info-box-text">Total Review/Testing</span>
              <span className="info-box-number">
                {loading ? (
                  <Spinner animation="border" size="sm" />
                ) : getCardCategory.length > 0 ? (
                  getCardCategory[3].count
                ) : (
                  0
                )}
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
              <span className="info-box-text">Total Done</span>
              <span className="info-box-number">
                {loading ? (
                  <Spinner animation="border" size="sm" />
                ) : getCardCategory.length > 0 ? (
                  getCardCategory[2].count
                ) : (
                  0
                )}
              </span>
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
            <div className="card-body pt-0 pb-0">
              {/* {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                0
              )} */}
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <>
                  <div className="row">
                    {/* /.col */}
                    <div className="col-md-12">
                      <div className="chart">
                        <div className="chartjs-size-monitor">
                          <div className="chartjs-size-monitor-expand">
                            <div className />
                          </div>
                          <div className="chartjs-size-monitor-shrink">
                            <div className />
                          </div>
                        </div>
                        {getDailyCountTask ? (
                          <ReactECharts option={getDailyCountTask} />
                        ) : (
                          ""
                        )}
                      </div>
                    </div>
                    {/* /.col */}
                  </div>
                  {/* /.row */}
                </>
              )}
            </div>
            {/* ./card-body */}
          </div>
          {/* /.card */}
        </div>
        {/* /.col */}
      </div>
    </>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(ProgressData);
