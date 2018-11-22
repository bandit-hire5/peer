package main

import (
	. "github.com/bandit/peer/config"
)

func main() {
	var config = Config{}
	config.Read()

	app := NewApp(config)
	app.Peer()
}
