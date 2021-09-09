package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hinha/PAM-Trello/app"
	"net/http"
)

func main() {
	err := connect("ws://localhost:8080/dashboard/inbox/ws?key=MzU3MjNjOWUwNjFhMjk3MzFmYjIyYjk5MjNmMTRlZWE5YTI0Y2E0MmQyNzk2M2ZhNDA1ZTQxZjgzOTRkYjhiODUyMWRjNmE4MzU3NzdhOGYzY2Y5Y2RjNDM3OTM5NjM3MjQ4YjRkNDE4OTMwNTEzNzM2Y2NlOWIxOWQxNmIxYTVlYWVmM2EzODEzODkwZWQzZmIwN2Q1Nzk2M2I2MTgwOGI0NDA1Y2IxNDZjMzdiOWNlM2NjNTJlYjM5MzFjYjEz")
	if err != nil {
		panic(err)
	}
}

func connect(url string) error {
	headers := make(http.Header)
	headers.Add("Origin", "http://192.168.1.2:4000")

	dialer := websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig:&tls.Config{
			InsecureSkipVerify: true,
		},
	}

	ws, _, err := dialer.Dial(url, headers)
	if err != nil {
		return err
	}

	go readConsole(ws)

	ws.Close()

	return nil
}

func readConsole(ws *websocket.Conn) {
	data := app.SocketEventStruct{
		EventItem: "performance",
		EventName: "update",
		EventPayload: map[string]string{
			"coba": "ping",
		},
	}

	by, _ := json.Marshal(data)
	for {
		err := ws.WriteMessage(websocket.BinaryMessage, by)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
