package main

import (
	"chat-app/internal/websocket"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTTP Server is running! Use /ws for WebSocket."))
	})

	// "/ws" 경로로 오는 웹소켓 요청을 internal/websocket 패키지의 ServeWs 함수가 처리하도록 합니다.
	http.HandleFunc("/ws", websocket.ServeWs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
