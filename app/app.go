package app

import (
	"github.com/bandit/blockchain-core"
	c "github.com/bandit/peer/config"
)

type App struct {
	config c.Config
	Ledger *core.Ledger
}

func NewApp(config c.Config) *App {
	app := &App{
		config: config,
	}

	app.initLedger()

	return app
}

func (a *App) Peer() {
	server := NewServer(a)
	server.Listen()
}

func (a *App) initLedger() {
	ledger := core.NewLedger(a.config.LedgerPath)
	ledger.CreateEmpty()

	a.Ledger = ledger
}
