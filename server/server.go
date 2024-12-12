package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"ws-echo-server/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// addr - server info
func Start(addr string) {
	http.HandleFunc("/ws", handleWebSoket)
	utils.Log(fmt.Sprintf("WebSocket server is running on %s...", addr))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		utils.Log(fmt.Sprintf("Server error: %v", err))
	}
}
