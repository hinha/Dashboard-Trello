import { connect as connectRedux } from "react-redux";
import { ListAccountModal } from "./Settings";
import { EDIT_USER_SETTING } from "../../services/setting";

// import * as ActionTypes from "../../store/actions";

const ModalManager = (props) => {
  const {
    closeFn = null,
    modal = "",
    onUpdateUserAccount = () => null,
    properties,
    dashboard,
  } = props;

  if (modal === EDIT_USER_SETTING) {
    let filter = dashboard.accounts.filter((item) => item.id === properties);
    return (
      <>
        <ListAccountModal
          closeFn={closeFn}
          dispatch={onUpdateUserAccount}
          data={filter.length > 0 ? filter[0] : {}}
          open={modal === EDIT_USER_SETTING}
        />
      </>
    );
  }

  return <></>;
};

const mapStateToProps = (state) => ({
  user: state.auth.currentUser,
  dashboard: state.dashboard.settings,
});

// const mapDispatchToProps = (dispatch) => ({
//   onUserLoad: (user) =>
//     dispatch({ type: ActionTypes.LOAD_USER_FORM, settings: user }),
// });

export default connectRedux(mapStateToProps, null)(ModalManager);
