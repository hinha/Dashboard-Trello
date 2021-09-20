import React, { useState, useEffect } from "react";
import { connect } from "react-redux";
import DataTable from "react-data-table-component";
import Spinner from "react-bootstrap/Spinner";

import { getClusters } from "../../services/analytics";
import { ClusterColumn, ClusterCoefficient } from "./cluster_data";

import ReactECharts from "echarts-for-react";
import * as echarts from "echarts";
import { transform } from "echarts-stat";

echarts.registerTransform(transform.clustering);

const Cluster = ({ credentials, token }) => {
  const [getClusterTable, setClusterTable] = useState([]);
  const [getScatterCluster, setScatterCluster] = useState({});
  const [getClusterCoefficient, setClusterCoefficient] = useState({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    let mounted = true;

    const fetchAPI = async () => {
      setLoading(true);
      const result = await getClusters(token, credentials, "2021");

      let newItem = [];
      result.card.forEach((item, index) => {
        let unf_check_items_complete = 0;
        let unf_count_check_items = 0;
        let created_at = item.created_at;
        let unf_comment = 0;
        if (item.check_items_complete !== undefined) {
          unf_check_items_complete = item.check_items_complete;
        }
        if (item.count_check_items !== undefined) {
          unf_count_check_items = item.count_check_items;
        }
        if (item.comment_count !== undefined) {
          unf_comment = item.comment_count;
        }

        let unf_cluster_1 = "-";
        if (item.cluster_1 !== undefined) {
          unf_cluster_1 = item.cluster_1;
        }
        let unf_cluster_2 = "-";
        if (item.cluster_2 !== undefined) {
          unf_cluster_2 = item.cluster_2;
        }
        let unf_cluster_3 = "-";
        if (item.cluster_3 !== undefined) {
          unf_cluster_3 = item.cluster_3;
        }
        let unf_cluster_4 = "-";
        if (item.cluster_4 !== undefined) {
          unf_cluster_4 = item.cluster_4;
        }
        let unf_cluster_5 = "-";
        if (item.cluster_5 !== undefined) {
          unf_cluster_5 = item.cluster_5;
        }
        let unf_cluster_6 = "-";
        if (item.cluster_6 !== undefined) {
          unf_cluster_6 = item.cluster_6;
        }
        let unf_cluster_7 = "-";
        if (item.cluster_7 !== undefined) {
          unf_cluster_7 = item.cluster_7;
        }
        let unf_cluster_8 = "-";
        if (item.cluster_8 !== undefined) {
          unf_cluster_8 = item.cluster_8;
        }
        let unf_cluster_9 = "-";
        if (item.cluster_9 !== undefined) {
          unf_cluster_9 = item.cluster_9;
        }
        let unf_cluster_10 = "-";
        if (item.cluster_10 !== undefined) {
          unf_cluster_10 = item.cluster_10;
        }

        newItem.push({
          no: index + 1,
          created_at: created_at,
          complete: unf_check_items_complete,
          total: unf_count_check_items,
          comment: unf_comment,
          cluster_1: unf_cluster_1,
          cluster_2: unf_cluster_2,
          cluster_3: unf_cluster_3,
          cluster_4: unf_cluster_4,
          cluster_5: unf_cluster_5,
          cluster_6: unf_cluster_6,
          cluster_7: unf_cluster_7,
          cluster_8: unf_cluster_8,
          cluster_9: unf_cluster_9,
          cluster_10: unf_cluster_10,
        });
      });
      setClusterTable(newItem);
      setScatterCluster(result.scatterClustering);
      setClusterCoefficient(ClusterCoefficient(result.average));
      setLoading(false);
    };

    if (mounted) {
      fetchAPI();
      console.log("mounted clusters", credentials);
    }

    return () => (mounted = false);
  }, [credentials, token]);

  return (
    <>
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
                <ReactECharts option={getScatterCluster} />
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
    </>
  );
};

const mapStateToProps = (state) => ({
  credentials: state.auth.credentials,
  token: state.auth.token,
});

export default connect(mapStateToProps, null)(Cluster);
