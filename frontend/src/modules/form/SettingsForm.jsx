import { useState } from "react";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import * as Yup from "yup";
import {
  TRELLO_USER_SETTING,
  ADD_USER_SETTING,
  ROLE_USER_SETTING,
} from "../../services/setting";

const handleFormToast = (msg) => {
  if (msg.error) {
    toast.error(msg.error || "Failed");
  }
  if (msg.status == "success") {
    toast.success("Successfully");
  }
};

const printFormError = (formik, key) => {
  if (formik.touched[key] && formik.errors[key]) {
    return <div>{formik.errors[key]}</div>;
  }
  return null;
};

export const FormTrello = (props) => {
  const { handleRequestAPI, listOptions, dispatch } = props;
  const [isSubmitting, setSubmitting] = useState(false);

  const formik = useFormik({
    initialValues: {
      id: "",
      boardID: "",
      name: "",
      memberID: "",
    },
    validationSchema: Yup.object({
      id: Yup.string()
        .min(5, "must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required(),
      boardID: Yup.string()
        .min(3, "must be 3 characters or more")
        .max(50, "Must be 30 characters or less")
        .required(),
      name: Yup.string()
        .min(3, "must be 3 characters or more")
        .max(30, "Must be 30 characters or less")
        .required(),
      memberID: Yup.string()
        .min(5, "must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required(),
    }),

    onSubmit: async (values) => {
      let postData = {
        user_id: values.id,
        board_name: values.name,
        member_id: values.memberID,
        board_id: values.boardID,
      };
      setSubmitting(true);
      const result = await handleRequestAPI(TRELLO_USER_SETTING, postData);
      handleFormToast(result);
      dispatch(result);
      setTimeout(() => {
        setSubmitting(false);
      }, 1000);
    },
  });

  return (
    <>
      {/* form start */}
      <form onSubmit={formik.handleSubmit}>
        <div className="card-body">
          <div className="form-group">
            <label>Select User</label>
            <select
              className="form-control"
              style={{ width: "100%" }}
              onChange={formik.handleChange}
              value={formik.values.id}
              {...formik.getFieldProps("id")}
            >
              {listOptions &&
                listOptions.map((item, index) => {
                  // var x = (day == "yes") ? "Good Day!" : (day == "no") ? "Good Night!" : "";
                  let str =
                    item.value !== ""
                      ? `${item.value} - ${item.label}`
                      : `${item.label}`;
                  return (
                    <option key={index} value={item.value}>
                      {str}
                    </option>
                  );
                })}
            </select>
            {printFormError(formik, "id")}
          </div>
          <div className="form-group">
            <label>Board Name</label>
            <input
              type="text"
              className="form-control"
              placeholder="Enter Name"
              {...formik.getFieldProps("name")}
            />
            {printFormError(formik, "name")}
          </div>
          <div className="form-group">
            <label>Board ID</label>
            <input
              type="text"
              className="form-control"
              placeholder="Enter ID"
              {...formik.getFieldProps("boardID")}
            />
            {printFormError(formik, "boardID")}
          </div>
          <div className="form-group">
            <label>Member ID</label>
            <input
              type="text"
              className="form-control"
              placeholder="Enter Id"
              {...formik.getFieldProps("memberID")}
            />
            {printFormError(formik, "memberID")}
          </div>
        </div>
        {/* /.card-body */}
        <div className="card-footer">
          <button
            type="submit"
            className="btn btn-primary"
            disabled={isSubmitting}
          >
            Create
          </button>
        </div>
      </form>
      {/* form end */}
    </>
  );
};

export const FormUser = (props) => {
  const { handleRequestAPI, dispatch } = props;
  const [isSubmitting, setSubmitting] = useState(false);

  const formik = useFormik({
    initialValues: {
      name: "",
      username: "",
      email: "",
      password: "",
    },
    validationSchema: Yup.object({
      name: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(50, "Must be 50 characters or less")
        .required(),
      username: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(30, "Must be 30 characters or less")
        .required(),
      email: Yup.string()
        .email()
        .min(5, "Must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required("Required"),
      password: Yup.string()
        .min(5, "Must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required("Required"),
    }),
    onSubmit: async (values) => {
      let postData = {
        name: values.name,
        username: values.username,
        password: values.password,
        email: values.email,
      };

      const result = await handleRequestAPI(ADD_USER_SETTING, postData);
      handleFormToast(result);
      dispatch(result);
      setTimeout(() => {
        setSubmitting(false);
      }, 1000);
    },
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <div className="card-body">
        <div className="form-group">
          <label>Full Name</label>
          <input
            type="text"
            className="form-control"
            placeholder="Enter Full Name"
            {...formik.getFieldProps("name")}
          />
          {printFormError(formik, "name")}
        </div>
        <div className="form-group">
          <label>Username</label>
          <input
            type="text"
            className="form-control"
            placeholder="Enter Username"
            {...formik.getFieldProps("username")}
          />
          {printFormError(formik, "username")}
        </div>
        <div className="form-group">
          <label>Email address</label>
          <input
            type="email"
            className="form-control"
            placeholder="Enter email"
            {...formik.getFieldProps("email")}
          />
          {printFormError(formik, "email")}
        </div>
        <div className="form-group">
          <label>Password</label>
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            {...formik.getFieldProps("password")}
          />
          {printFormError(formik, "password")}
        </div>
      </div>
      {/* /.card-body */}
      <div className="card-footer">
        <button
          type="submit"
          className="btn btn-primary"
          disabled={isSubmitting}
        >
          Submit
        </button>
      </div>
    </form>
  );
};

export const FormRole = (props) => {
  const { handleRequestAPI, listOptions, listAccount } = props;
  const [isSubmitting, setSubmitting] = useState(false);

  const formik = useFormik({
    initialValues: {
      user: "",
      role: "",
      permission: "",
    },
    validationSchema: Yup.object({
      user: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(50, "Must be 50 characters or less")
        .required(),
      role: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(50, "Must be 50 characters or less")
        .required(),
      permission: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(50, "Must be 50 characters or less")
        .required(),
    }),
    onSubmit: async (values) => {
      let postData = {
        userID: values.user,
        role: values.role,
        permission: values.permission,
      };

      const result = await handleRequestAPI(ROLE_USER_SETTING, postData);
      handleFormToast(result);
      setTimeout(() => {
        setSubmitting(false);
      }, 1000);
    },
  });

  return (
    <form onSubmit={formik.handleSubmit}>
      <div className="card-body">
        <div className="form-group">
          <label>Select User</label>
          <select
            className="form-control select2"
            style={{ width: "100%" }}
            onChange={formik.handleChange}
            value={formik.values.user}
            {...formik.getFieldProps("user")}
          >
            <option value="" selected={true}>
              User
            </option>
            {listAccount &&
              listAccount.map((item, index) => {
                return (
                  <option key={index} value={item.id}>
                    {item.name}
                  </option>
                );
              })}
          </select>
          {printFormError(formik, "user")}
        </div>
        <div className="form-group">
          <label>Role</label>
          <select
            className="form-control select2"
            style={{ width: "100%" }}
            onChange={formik.handleChange}
            value={formik.values.role}
            {...formik.getFieldProps("role")}
          >
            <option value="" selected={true}>
              Select Role
            </option>
            {listOptions.authority &&
              listOptions.authority.map((item, index) => {
                return (
                  <option key={index} value={item.name}>
                    {item.name}
                  </option>
                );
              })}
          </select>
        </div>
        <div className="form-group">
          <label>Permission</label>
          <select
            className="form-control select2"
            style={{ width: "100%" }}
            onChange={formik.handleChange}
            value={formik.values.permission}
            {...formik.getFieldProps("permission")}
          >
            <option value="" selected={true}>
              Select Permission
            </option>
            {listOptions.permission &&
              listOptions.permission.map((item, index) => {
                return (
                  <option key={index} value={item.name}>
                    {item.name}
                  </option>
                );
              })}
          </select>
        </div>
      </div>
      {/* /.card-body */}
      <div className="card-footer">
        <button
          type="submit"
          className="btn btn-primary"
          disabled={isSubmitting}
        >
          Create
        </button>
      </div>
    </form>
  );
};
