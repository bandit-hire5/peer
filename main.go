package main

import (
	. "github.com/bandit/peer/config"

	"github.com/bandit/peer/app"
)

func main() {
	var config = Config{}
	config.Read()

	app := app.NewApp(config)
	app.Peer()
}
