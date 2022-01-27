import React, { useRef, useEffect, useState } from "react";
import { useHistory, Link } from "react-router-dom";
import { connect } from "react-redux";
import { DateTime } from "luxon";
import { useTranslation } from "react-i18next";

import * as ActionTypes from "../../../../store/actions";

const UserDropdown = ({ user, onUserLogout }) => {
  const dropdownRef = useRef(null);
  const history = useHistory();
  const [t] = useTranslation();

  const [dropdownState, updateDropdownState] = useState({
    isDropdownOpen: false,
  });

  const toggleDropdown = () => {
    updateDropdownState({ isDropdownOpen: !dropdownState.isDropdownOpen });
  };

  const handleClickOutside = (event) => {
    if (
      dropdownRef &&
      dropdownRef.current &&
      !dropdownRef.current.contains(event.target)
    ) {
      updateDropdownState({ isDropdownOpen: false });
    }
  };

  const logOut = (event) => {
    toggleDropdown();
    event.preventDefault();
    // TODO: Need call logout api
    onUserLogout();
    history.push("/");
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleClickOutside, false);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside, false);
    };
  });

  let className = "dropdown-menu dropdown-menu-lg dropdown-menu-right";

  if (dropdownState.isDropdownOpen) {
    className += " show";
  }

  return (
    <li ref={dropdownRef} className="nav-item dropdown user-menu">
      <a
        onClick={toggleDropdown}
        type="button"
        className="nav-link dropdown-toggle"
        data-toggle="dropdown"
      >
        <img
          src={user.picture || "../dist/img/user2-160x160.jpg"}
          className="user-image img-circle elevation-2"
          alt="User"
        />
      </a>
      <ul className={className}>
        <li className="user-header bg-primary">
          <img
            src={user.picture || "../dist/img/user2-160x160.jpg"}
            className="img-circle elevation-2"
            alt="User"
          />
          <p>
            {user.email}
            <small>
              <span>Member since </span>
              <span>
                {DateTime.fromISO(user.createdAt).toFormat("dd mm yyyy")}
              </span>
            </small>
          </p>
        </li>

        <li className="user-footer">
          <Link
            to="/settings"
            onClick={toggleDropdown}
            className="btn btn-default btn-flat"
          >
            {t("header.user.profile")}
          </Link>
          <button
            type="button"
            className="btn btn-default btn-flat float-right"
            onClick={logOut}
          >
            {t("login.button.signOut")}
          </button>
        </li>
      </ul>
    </li>
  );
};

const mapStateToProps = (state) => ({
  user: state.auth.currentUser,
});

const mapDispatchToProps = (dispatch) => ({
  onUserLogout: () => dispatch({ type: ActionTypes.LOGOUT_USER }),
});

export default connect(mapStateToProps, mapDispatchToProps)(UserDropdown);
