const domElement = document.querySelector("#table-container");

class App extends React.Component {
    constructor() {
        super();
        this.state = {
            chatUserList: [],
            message: null,
            selectedUserID: null,
            userID: null
        }
        this.webSocketConnection = null;
    }

    componentDidMount() {
        this.setWebSocketConnection();
        this.subscribeToSocketMessage();
    }

    setWebSocketConnection() {
        if (window["WebSocket"]) {
            this.webSocketConnection = new WebSocket("ws://" + document.location.host + "/dashboard/ws");
        }
    }

    subscribeToSocketMessage = () => {
        if (this.webSocketConnection === null) {
            return;
        }

        this.webSocketConnection.onclose = (evt) => {
            this.setState({
                message: 'Your Connection is closed.',
                chatUserList: []
            });
        };

        this.webSocketConnection.onmessage = (event) => {
            try {
                const socketPayload = JSON.parse(event.data);
                switch (socketPayload.eventName) {
                    case 'join':
                    case 'disconnect':
                        if (!socketPayload.eventPayload) {
                            return
                        }

                        const userInitPayload = socketPayload.eventPayload;

                        this.setState({
                            chatUserList: userInitPayload.users,
                            userID: this.state.userID === null ? userInitPayload.userID : this.state.userID
                        });

                        break;

                    case 'message response':

                        if (!socketPayload.eventPayload) {
                            return
                        }

                        const messageContent = socketPayload.eventPayload;
                        const sentBy = messageContent.username ? messageContent.username : 'An unnamed fellow'
                        const actualMessage = messageContent.message;

                        this.setState({
                            message: `${sentBy} says: ${actualMessage}`
                        });

                        break;

                    default:
                        break;
                }
            } catch (error) {
                console.log(error)
                console.warn('Something went wrong while decoding the Message Payload')
            }
        };

    };

    setNewUserToChat = (event) => {
        if (event.target && event.target.value) {
            if (event.target.value === "select-user") {
                alert("Select a user to chat");
                return;
            }
            this.setState({
                selectedUserID: event.target.value
            })
        }
    };

    getChatList() {
        if (this.state.chatUserList.length === 0) {
            return(
                <table className="table m-0">
                    <thead>
                    <tr>
                        <th>Name</th>
                        <th>Status</th>
                        <th>Title</th>
                    </tr>
                    </thead>
                    <tbody>
                    </tbody>
                </table>
            )
        }
        return (
            <table className="table m-0">
                <thead>
                <tr>
                    <th>Name</th>
                    <th>Status</th>
                    <th>Title</th>
                </tr>
                </thead>
                <tbody>
                {
                    this.state.chatUserList.map(user => {
                        if (user.userID !== this.state.userID) {
                            return (
                                <tr>
                                    <td>{user.name}</td>
                                    <td><span className="badge badge-success">Online</span></td>
                                    <td>
                                        <div className="sparkbar" data-color="#00a65a" data-height="20">Backend</div>
                                    </td>
                                </tr>
                            )
                        }
                    })
                }

                </tbody>
            </table>
        );
    }

    render() {
        return (
            <React.Fragment>
                {this.getChatList()}
            </React.Fragment>
        );
    }
}


ReactDOM.render(<App />, domElement)