package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/DmitriiTrifonov/octo-giggle/backend/internal/model"
	"github.com/gorilla/websocket"
)

type WS struct {
	ctx   context.Context
	Input chan model.Request
}

func NewWS(ctx context.Context, input chan model.Request) *WS {
	return &WS{
		ctx:   ctx,
		Input: input,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *WS) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(rw, request, nil)
	log.Println("opening websocket")
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	ch := make(chan int, 0)
	go readMessage(ch, conn)

	for {
		select {
		case <-ch:
			log.Println("closing websocket")
			return
		case <-ws.ctx.Done():
			return
		case data := <-ws.Input:
			err = conn.WriteJSON(data)
			if err != nil {
				log.Println(err)
				http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}
}

func readMessage(ch chan int, conn *websocket.Conn) {
	for {
		mt, _, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			ch <- mt
			return
		}
	}
}
