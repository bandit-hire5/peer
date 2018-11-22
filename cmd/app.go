package main

import (
	"log"
	"net/http"

	"github.com/bandit/blockchain-core"
	"github.com/bandit/peer/config"
	"github.com/bandit/peer/rest/routes"
	ws "github.com/bandit/peer/websocket"
)

type App struct {
	config config.Config
	ledger *core.Ledger
}

func NewApp(cfg config.Config) *App {
	app := &App{
		config: cfg,
	}

	app.initLedger()

	return app
}

func (self *App) Peer() {
	server := ws.NewServer()
	server.Config = &self.config
	server.Ledger = self.ledger

	go server.Listen()

	router := routes.Router(server)

	log.Fatal(http.ListenAndServe(":"+self.config.Port, router))
}

func (self *App) initLedger() {
	ledger := core.NewLedger(self.config.LedgerPath)

	err := ledger.CreateEmpty()
	if err != nil {
		panic(err)
	}

	self.ledger = ledger
}
