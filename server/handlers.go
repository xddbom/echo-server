package server

import (
	"io"
	"net/http"

	"ws-echo-server/utils"
)

func handleWebSoket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Log("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	utils.Log("New WebSocket connection establish")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				utils.Log("Client disconnected")
			} else {
				utils.Log("Read error:", err)
			}
			break
		}

		utils.Log("Received message:", string(msg))

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			utils.Log("Write error:", err)
			break
		}
	}
}
