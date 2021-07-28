export default class Socket {
  constructor(userId, credentials) {
    this.userID = userId;
    this.credentials = credentials;
    this.state = {
      performance: {
        todoItem: 0,
        doneItem: 0,
      },
    };
    this.webSocketConnection = null;
  }

  setWebsocket() {
    this.webSocketConnection = new WebSocket(
      "ws://localhost:8080/dashboard/inbox/ws?key=" + this.credentials
    );
  }

  sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  keepAlive() {
    this.setWebsocket();
    if (this.webSocketConnection !== null) {
      this.subscribeToSocketMessage();
    }
    this.sleep(10000);
  }

  subscribeToSocketMessage() {
    this.setWebsocket();

    if (this.webSocketConnection === null) {
      return;
    }

    this.webSocketConnection.onclose = () => {
      console.log("Your Connection is closed.");
      var that = this;
      setTimeout(function () {
        that.subscribeToSocketMessage();
      }, 1000);
    };

    this.webSocketConnection.onopen = () => {
      this.state.performance = {
        doneItem: 1,
        todoItem: 2,
      };
    };

    this.webSocketConnection.onmessage = (event) => {
      try {
        const socketPayload = JSON.parse(event.data);
        switch (socketPayload.eventName) {
          case "response":
            if (socketPayload.eventItem === "performance") {
              this.state.performance = {
                doneItem: 1,
                todoItem: 2,
              };
            }
            break;
        }
      } catch (e) {
        console.error(e);
      }
    };
  }

  tesKiremClick(state) {
    if (this.webSocketConnection === null) {
      return {};
    }

    if (this.webSocketConnection.readyState === 0) {
      console.log(this.webSocketConnection.readyState, "ready");
      var that = this;
      setTimeout(function () {
        that.tesKiremClick(state);
      }, 1000);
    } else {
      this.webSocketConnection.send(
        JSON.stringify({
          eventItem: state,
          eventName: "update",
          eventPayload: {
            userID: this.userID,
            message: "ping dari browser",
          },
        })
      );
    }

    return this.state.performance;
  }
}
