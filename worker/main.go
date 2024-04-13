package main

import (
	"shared/config"
	db "shared/db/sqlc"
	"shared/events"
	"sync"
)

// func runTaskProcessor(redisOpt asynq.RedisClientOpt, store){

// }

func main() {
	config.Set()
	db.Connect()
	var wg sync.WaitGroup
	wg.Add(1)
	go events.InitTaskProcessor(*db.DbStore)
	wg.Wait()
}
