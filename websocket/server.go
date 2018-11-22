package websocket

import (
	"encoding/json"
	"io"
	"log"

	"github.com/bandit/blockchain-core"
	"github.com/bandit/peer/config"

	"golang.org/x/net/websocket"
)

type Server struct {
	Config *config.Config
	Ledger *core.Ledger
	ws     *websocket.Conn
	mCh    chan *core.Message
	doneCh chan bool
	errCh  chan error
}

func NewServer() *Server {
	mCh := make(chan *core.Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		mCh:    mCh,
		doneCh: doneCh,
		errCh:  errCh,
	}
}

func (server *Server) SendMessage(message *core.Message) error {
	err := websocket.JSON.Send(server.ws, message)
	if err != nil {
		return err
	}

	return nil
}

func (self *Server) Listen() {
	log.Println("Listening server...")

	ws, err := websocket.Dial(self.Config.Ws, "", self.Config.Origin)
	if err != nil {
		panic(err)
	}

	self.ws = ws

	go self.listenRead(ws)

	for {
		select {
		case msg := <-self.mCh:
			log.Println("Message: ", msg)

			if msg.Type == "block" {
				var block core.Block

				json.Unmarshal([]byte(msg.Body), &block)

				self.Ledger.AddBlock(&block)
			}
		}
	}
}

func (self *Server) listenRead(ws *websocket.Conn) {
	for {
		var msg *core.Message

		err := websocket.JSON.Receive(ws, &msg)

		if err == io.EOF {
			self.doneCh <- true
		} else if err != nil {
			self.errCh <- err
		} else {
			self.mCh <- msg
		}
	}
}
