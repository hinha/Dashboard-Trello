import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import _ from "lodash";

import { UPDATE_ANALYTIC_TRELLO } from "../services/analytics";

const TrelloBoard = ({ onClickSidebarApi }) => {
  const [getCardTodo, setCardTodo] = useState([]);
  const [getCardProgress, setCardProgress] = useState([]);
  const [getCardDone, setCardDone] = useState([]);

  useEffect(() => {
    let mounted = true;
    const fetchAPI = async () => {
      const result = await onClickSidebarApi(UPDATE_ANALYTIC_TRELLO);
      var grouped = _.mapValues(
        _.groupBy(result.data, "card_category"),
        (clist) => clist.map((car) => _.omit(car, "make"))
      );

      setCardTodo(grouped.TODO);
      setCardProgress(grouped.ON_PROGRESS);
      setCardDone(grouped.DONE);
    };

    if (mounted) {
      fetchAPI();
    }
    return () => (mounted = false);
  }, [onClickSidebarApi]);
  return (
    <div className="content-wrapper kanban">
      <section className="content-header">
        <div className="container-fluid">
          <div className="row">
            <div className="col-sm-6">
              <h1>Trello Board</h1>
            </div>
            <div className="col-sm-6 d-none d-sm-block">
              <ol className="breadcrumb float-sm-right">
                <li className="breadcrumb-item">
                  <Link to="/dashboard">Home</Link>
                </li>
                <li className="breadcrumb-item active">Trello Board</li>
              </ol>
            </div>
          </div>
        </div>
      </section>
      <section className="content pb-3">
        <div className="container-fluid h-100">
          <div className="card card-row card-secondary">
            <div className="card-header">
              <h3 className="card-title">TO DO</h3>
            </div>
            <div className="card-body">
              {getCardTodo &&
                getCardTodo.map((val, index) => {
                  return (
                    <div className="card card-primary card-outline" key={index}>
                      <div className="card-header">
                        <h5 className="card-title">{val.card_name}</h5>
                        <div className="card-tools">
                          <a
                            className="btn btn-tool btn-link"
                            onClick={() => window.open(val.card_url)}
                          >
                            #{val.ID}
                          </a>
                          <a href="#" className="btn btn-tool">
                            <i className="fas fa-pen" />
                          </a>
                        </div>
                      </div>
                      <div className="card-body">
                        <p>Assingne: {val.card_member_name}</p>
                      </div>
                    </div>
                  );
                })}
            </div>
          </div>
          <div className="card card-row card-primary">
            <div className="card-header">
              <h3 className="card-title">In Progress</h3>
            </div>
            <div className="card-body">
              {getCardProgress &&
                getCardProgress.map((val, index) => {
                  return (
                    <div className="card card-primary card-outline" key={index}>
                      <div className="card-header">
                        <h5 className="card-title">{val.card_name}</h5>
                        <div className="card-tools">
                          {/* <Link to="route" target={val.card_url} rel="noopener noreferrer" className="btn btn-tool btn-link"/> */}
                          <a
                            className="btn btn-tool btn-link"
                            onClick={() => window.open(val.card_url)}
                          >
                            #{val.ID}
                          </a>
                          <a href="#" className="btn btn-tool">
                            <i className="fas fa-pen" />
                          </a>
                        </div>
                      </div>
                      <div className="card-body">
                        <p>Assingne: {val.card_member_name}</p>
                      </div>
                    </div>
                  );
                })}
            </div>
          </div>
          <div className="card card-row card-success">
            <div className="card-header">
              <h3 className="card-title">Done</h3>
            </div>
            <div className="card-body">
              {getCardDone &&
                getCardDone.map((val, index) => {
                  return (
                    <div className="card card-primary card-outline" key={index}>
                      <div className="card-header">
                        <h5 className="card-title">{val.card_name}</h5>
                        <div className="card-tools">
                          {/* <Link to="route" target={val.card_url} rel="noopener noreferrer" className="btn btn-tool btn-link"/> */}
                          <a
                            className="btn btn-tool btn-link"
                            onClick={() => window.open(val.card_url)}
                          >
                            #{val.ID}
                          </a>
                          <a href="#" className="btn btn-tool">
                            <i className="fas fa-pen" />
                          </a>
                        </div>
                      </div>
                      <div className="card-body">
                        <p>Assingne: {val.card_member_name}</p>
                      </div>
                    </div>
                  );
                })}
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

export default connect(mapStateToProps, null)(TrelloBoard);
