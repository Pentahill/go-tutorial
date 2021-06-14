package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func recieve(conn *websocket.Conn, reCh chan string, seCh chan string) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("receive error")
		}

		if msg != nil {
			fmt.Println("receive msg ", string(msg))
			reCh <- string(msg)
		}
	}
}

func send(conn *websocket.Conn, reCh chan string, seCh chan string) {
	for {
		select {
		case msg := <-reCh:
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
			fmt.Println("send message", msg)
		}
	}
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err1 := upgrader.Upgrade(w, r, nil)
	if err1 != nil {
		fmt.Println(err1)
		w.Write([]byte("upgrade error"))
		return
	}

	writer, err2 := conn.NextWriter(websocket.TextMessage)
	if err2 != nil {
		writer.Write([]byte("get writer error"))
		return
	}

	receiveCh := make(chan string, 10)
	sendCh := make(chan string)

	go recieve(conn, receiveCh, sendCh)
	go send(conn, receiveCh, sendCh)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func main() {
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		handleWS(w, r)
// 	})
// 	err := http.ListenAndServe(":8000", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
