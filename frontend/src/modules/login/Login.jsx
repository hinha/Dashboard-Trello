import React, { useState } from "react";
import { Link, useHistory } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import { connect } from "react-redux";
import { useFormik } from "formik";
import { useTranslation } from "react-i18next";
import "react-toastify/dist/ReactToastify.css";

import * as Yup from "yup";

import * as AuthService from "./../../services/auth";
import Button from "./../../components/button/button";
import * as ActionTypes from "../../store/actions";

const Login = ({ onUserLogin }) => {
  const [isAuthLoading, setAuthLoading] = useState(false);

  const history = useHistory();
  const [t] = useTranslation();

  const login = async (username, password, remember_me) => {
    try {
      setAuthLoading(true);
      await AuthService.getToken();

      const token = await AuthService.loginByAuth(
        username,
        password,
        remember_me
      );
      toast.success("Login is succeed!");
      setAuthLoading(false);
      onUserLogin(token.data.token);
      history.push("/dashboard");
    } catch (error) {
      setAuthLoading(false);
      toast.error(
        (error.response && error.response.data && error.response.data.error) ||
          "Failed"
      );
    }
  };

  const printFormError = (formik, key) => {
    if (formik.touched[key] && formik.errors[key]) {
      return <div>{formik.errors[key]}</div>;
    }
    return null;
  };

  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
      remember_me: false,
    },
    validationSchema: Yup.object({
      username: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(30, "Must be 30 characters or less")
        .required("Required"),
      password: Yup.string()
        .min(5, "Must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required("Required"),
    }),
    onSubmit: (values) => {
      login(values.username, values.password, values.remember_me);
    },
  });

  document.getElementById("root").classList = "hold-transition login-page";

  return (
    <div className="login-box">
      <div className="card card-outline card-primary">
        <div className="card-header text-center">
          <Link to="/" className="h1">
            <b>Admin</b>
            <span>LTE</span>
          </Link>
        </div>
        <div className="card-body">
          <p className="login-box-msg">{t("login.label.signIn")}</p>
          <ToastContainer />
          <form onSubmit={formik.handleSubmit}>
            <div className="mb-3">
              <div className="input-group">
                <input
                  type="text"
                  className="form-control"
                  placeholder="username"
                  {...formik.getFieldProps("username")}
                />
                <div className="input-group-append">
                  <div className="input-group-text">
                    <span className="fas fa-envelope" />
                  </div>
                </div>
              </div>
              {formik.touched.username && formik.errors.username ? (
                <div>{formik.errors.username}</div>
              ) : null}
            </div>
            <div className="mb-3">
              <div className="input-group">
                <input
                  type="password"
                  className="form-control"
                  placeholder="Password"
                  {...formik.getFieldProps("password")}
                />
                <div className="input-group-append">
                  <div className="input-group-text">
                    <span className="fas fa-lock" />
                  </div>
                </div>
              </div>
              {printFormError(formik, "password")}
            </div>
            <div className="row">
              <div className="col-8">
                <div className="icheck-primary">
                  <input
                    type="checkbox"
                    id="remember"
                    name="remember_me"
                    {...formik.getFieldProps("remember_me")}
                  />
                  <label htmlFor="remember">
                    {t("login.label.rememberMe")}
                  </label>
                </div>
              </div>
              <div className="col-4">
                <Button block type="submit" isLoading={isAuthLoading}>
                  {t("login.button.signIn.label")}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

const mapDispatchToProps = (dispatch) => ({
  onUserLogin: (token) => dispatch({ type: ActionTypes.LOGIN_USER, token }),
});

export default connect(null, mapDispatchToProps)(Login);
