package main

import (
	"os"
	"shared/config"
	db "shared/db/sqlc"
	"shared/events"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config.Set()
	db.Connect()
	var wg sync.WaitGroup
	wg.Add(1)
	go events.InitTaskProcessor(*db.DbStore, true)
	wg.Wait()
}
