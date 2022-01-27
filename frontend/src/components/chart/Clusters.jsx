import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import DataTable from "react-data-table-component";
import Spinner from "react-bootstrap/Spinner";

import { getClusters } from "../../services/analytics";
import { ClusterColumn } from "./cluster_data";
import { ResultAnalysisColumn } from "./analysis_result";

import ReactECharts from "echarts-for-react";

const Cluster = ({ credentials, token }) => {
  const [getWeight, setWeight] = useState([]);
  const [getClusterTable, setClusterTable] = useState([]);
  const [getClusterPlot, setClusterPlot] = useState({});
  const [getClusterCoefficient, setClusterCoefficient] = useState({});
  const [getActivityPlot, setActivityPlot] = useState({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    let mounted = true;

    const fetchAPI = async () => {
      setLoading(true);
      const result = await getClusters(token, credentials, "2021");
      console.log(result);

      let newItem = [];
      result.card.forEach((item, index) => {
        let unf_check_items_complete = 0;
        let unf_active = 0;
        let unf_years = 0;
        let unf_distance = 0;
        let unf_structure = 0;
        let time = item.time;
        let years = item.years;
        if (item.check_items_complete !== undefined) {
          unf_check_items_complete = item.check_items_complete;
        }
        if (item.active !== undefined) {
          unf_active = item.active;
        }
        if (item.years !== undefined) {
          unf_years = years;
        }
        if (item.kedekatan !== undefined) {
          unf_distance = item.kedekatan.toFixed(3);
        }
        if (item.struktur !== undefined) {
          unf_structure = item.struktur;
        }

        let unf_cluster_1 = "-";
        if (item.cluster_1 !== undefined) {
          unf_cluster_1 = item.cluster_1.toFixed(3);
        }
        let unf_cluster_2 = "-";
        if (item.cluster_2 !== undefined) {
          unf_cluster_2 = item.cluster_2.toFixed(3);
        }
        let unf_cluster_3 = "-";
        if (item.cluster_3 !== undefined) {
          unf_cluster_3 = item.cluster_3.toFixed(3);
        }
        let unf_cluster_4 = "-";
        if (item.cluster_4 !== undefined) {
          unf_cluster_4 = item.cluster_4.toFixed(3);
        }

        newItem.push({
          no: index + 1,
          time: time,
          complete: unf_check_items_complete,
          active: unf_active,
          years: unf_years,
          distance: unf_distance,
          structure: unf_structure,
          cluster_1: unf_cluster_1,
          cluster_2: unf_cluster_2,
          cluster_3: unf_cluster_3,
          cluster_4: unf_cluster_4,
        });
      });
      setWeight(result.weight);
      setClusterTable(newItem);
      setClusterPlot(result.scatter_clustering);
      setActivityPlot(result.activity);
      // setScatterCluster(result.scatterClustering);
      setClusterCoefficient(result.average_plot);
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
        <div className="col-md-12">
          <div className="card">
            <div className="card-header">
              <h5 className="card-title">Activity</h5>
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
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <ReactECharts option={getActivityPlot} />
              )}
            </div>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col-md-12">
          <div className="card">
            <div className="card-header">
              <h5 className="card-title">Cluster Data</h5>
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
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <DataTable
                  columns={ClusterColumn}
                  data={getClusterTable}
                  defaultSortFieldId={2}
                  pagination
                />
              )}
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
              <h5 className="card-title">Clustering</h5>
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
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <img src={getClusterPlot} />
              )}
            </div>
          </div>
        </div>
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
            <div className="card-body">
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <ReactECharts option={getClusterCoefficient} />
              )}
            </div>
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
            <div className="card-body pt-0 pb-0">
              {loading ? (
                <div className="w-100 d-flex justify-content-center align-items-center">
                  <Spinner animation="border" />
                </div>
              ) : (
                <DataTable
                  columns={ResultAnalysisColumn}
                  data={getWeight}
                  defaultSortFieldId={1}
                />
              )}
              {/* /.row */}
            </div>
            {/* ./card-body */}
          </div>
        </div>
      </div>
    </>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(Cluster);
