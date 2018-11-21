package app

import (
	"encoding/json"
	"io"
	"log"

	"github.com/bandit/blockchain-core"
	c "github.com/bandit/peer/config"

	"golang.org/x/net/websocket"
)

type Server struct {
	config *c.Config
	ledger *core.Ledger
	mCh    chan *Message
	doneCh chan bool
	errCh  chan error
}

func NewServer(a *App) *Server {
	mCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		&a.config,
		a.Ledger,
		mCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Listen() {
	log.Println("Listening server...")

	ws, err := websocket.Dial(s.config.Ws, "", s.config.Origin)
	if err != nil {
		panic(err)
	}

	go s.listenRead(ws)

	for {
		select {
		case msg := <-s.mCh:
			//log.Println("Message: ", msg)

			if msg.Type == "block" {
				var block core.Block

				json.Unmarshal([]byte(msg.Body), &block)

				s.ledger.AddBlock(&block)
			}
		}
	}
}

func (s *Server) listenRead(ws *websocket.Conn) {
	for {
		var msg *Message

		err := websocket.JSON.Receive(ws, &msg)

		if err == io.EOF {
			s.doneCh <- true
		} else if err != nil {
			s.errCh <- err
		} else {
			s.mCh <- msg
		}
		return
	}
}
