import React from "react";
import classes from "./PageLoading.module.scss";

const PageLoading = () => {
  return (
    <div className={classes.loading}>
      <span>L</span>
      <span>O</span>
      <span>A</span>
      <span>D</span>
      <span>I</span>
      <span>N</span>
      <span>G</span>
      <span>.</span>
      <span>.</span>
      <span>.</span>
    </div>
  );
};

export default PageLoading;
