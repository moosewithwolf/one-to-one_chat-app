package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader 는 HTTP 연결을 웹소켓 연결로 업그레이드하는 설정을 담당합니다.
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 개발 중에는 모든 출처를 허용합니다.
		// 실제 배포 시에는 보안을 위해 특정 출처만 허용하도록 수정해야 합니다.
		return true
	},
}
