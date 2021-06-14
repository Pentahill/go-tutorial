package connection

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan  chan []byte
	outChan chan []byte
	closeChan chan byte
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn := &Connection{
		wsConn: wsConn
		inChan: make(chan []byte, 1000)
		outChan: make(chan []byte, 1000)
		closeChan: make(chan byte)
	}

	go connection.readLoop()
	go connection.writeLoop()

	return conn, nil
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data <- conn.inChan:
	case <- conn.closeChan:
	}
	
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <- conn.closeChan:
	}
	return
}

var once sync.Once

func (conn *Connection) Close() (err error) {
	conn.Close()
	once.Do(func() {
		close(conn.closeChan)
	})
}

func (conn *Connection) readLoop() {
	var(
		data []byte
		err error
	)

	for{
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		
		select{
		case conn.inChan <- data:
		case <- conn.closeChan:
			goto ERR
		}
		
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var(
		data []byte
		err error
	)
	for{
		select{
		case data <- outChan:
			if err = conn.wsConn.WriteMessage(websocket.TextMessage, data): err != nil {
				goto ERR
			}
		case <-closeChan:
			goto ERR
		}
		
	}

ERR:
	conn.Close()
}
