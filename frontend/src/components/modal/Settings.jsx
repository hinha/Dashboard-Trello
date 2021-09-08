import { useState } from "react";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import { Alert } from "react-bootstrap";
import { useFormik } from "formik";
import * as Yup from "yup";

const printFormError = (formik, key) => {
  if (formik.touched[key] && formik.errors[key]) {
    return <div>{formik.errors[key]}</div>;
  }
  return null;
};

export const ListAccountModal = ({
  closeFn = () => null,
  dispatch = () => null,
  open = true,
  data = {},
}) => {
  const [isSubmitting, setSubmitting] = useState(false);
  const [showError, setShowError] = useState(false);
  const [getMsgError, setMsgError] = useState("");

  const formik = useFormik({
    initialValues: {
      name: data.name || "",
      email: data.email || "",
      activate: !data.suspend_status ? "activate" : "deactivate",
    },
    validationSchema: Yup.object({
      name: Yup.string()
        .min(3, "Must be 3 characters or more")
        .max(50, "Must be 50 characters or less")
        .required(),
      email: Yup.string()
        .email()
        .min(5, "Must be 5 characters or more")
        .max(30, "Must be 30 characters or less")
        .required("Required"),
      activate: Yup.string().min(1, "Must be 5 characters or more").required(),
    }),
    onSubmit: async (values) => {
      let isSuspend = values.activate === "activate" ? false : true;
      let postData = {
        id: data.id,
        name: values.name,
        email: values.email,
        suspend: isSuspend,
      };

      const result = await dispatch(postData);
      setSubmitting(true);

      if (result.error) {
        setShowError(true);
        setMsgError(result.error);
      }
      data.name = values.name;
      data.email = values.email;
      data.suspend_status = isSuspend;

      setTimeout(() => {
        setSubmitting(false);
      }, 1000);
      closeFn(true);
    },
  });

  let errLayout;
  if (showError) {
    errLayout = (
      <Alert
        className="mr-1 ml-1"
        variant="danger"
        onClose={() => setShowError(false)}
        dismissible
      >
        <p>{getMsgError}</p>
      </Alert>
    );
  }

  return (
    <Modal show={open} onHide={closeFn}>
      <Modal.Header>
        <Modal.Title>Update User</Modal.Title>
      </Modal.Header>
      {errLayout}

      <form onSubmit={formik.handleSubmit}>
        <Modal.Body>
          <div className="form-group">
            <label>Name</label>
            <input
              type="text"
              className="form-control"
              {...formik.getFieldProps("name")}
            />
            {printFormError(formik, "name")}
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
          {["radio"].map((type) => (
            <div
              key={`inline-${type}`}
              className="mb-3 pl-1"
              onChange={formik.handleChange}
              {...formik.getFieldProps("activate")}
            >
              <Form.Check
                inline
                label="Activate"
                name="activate"
                type={type}
                id={`inline-${type}-1`}
                defaultChecked={!data.suspend_status ? true : false}
                value="activate"
              />
              <Form.Check
                inline
                label="Deactivate"
                name="activate"
                type={type}
                id={`inline-${type}-2`}
                defaultChecked={data.suspend_status ? true : false}
                value="deactivate"
              />
            </div>
          ))}
          {printFormError(formik, "activate")}
        </Modal.Body>
        <Modal.Footer>
          <button type="button" onClick={closeFn} className="btn btn-secondary">
            Close
          </button>
          <button
            type="submit"
            className="btn btn-primary"
            disabled={isSubmitting}
          >
            Update
          </button>
        </Modal.Footer>
      </form>
    </Modal>
  );
};
