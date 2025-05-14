package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket" // gorilla/websocket 임포트 추가
)

// ServeWs 함수는 웹소켓 연결 요청을 처리합니다.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket endpoint hit by", r.RemoteAddr)

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	defer conn.Close()

	log.Println("Client successfully connected:", conn.RemoteAddr())

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error from %s: %v", conn.RemoteAddr(), err)
			} else {
				log.Printf("Read error from %s: %v", conn.RemoteAddr(), err)
			}
			break
		}

		log.Printf("Received message from %s: %s (type: %d)\n", conn.RemoteAddr(), string(p), messageType)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Write error to %s: %v", conn.RemoteAddr(), err)
			break
		}
	}
	log.Println("Client disconnected:", conn.RemoteAddr())
}
