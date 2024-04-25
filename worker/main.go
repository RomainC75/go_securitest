package main

import (
	"os"
	"shared/config"
	"shared/events"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config.Set()
	var wg sync.WaitGroup
	wg.Add(1)
	go events.InitTaskProcessor(true)
	wg.Wait()
}
