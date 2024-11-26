package kafka_helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	shared_dto "shared/dto"
	"shared/scenarios/selector"
	shared_utils "shared/utils"
	"syscall"
	"time"
)

func (kh *KafkaHandler) Listen() {
	// Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := kh.c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

			var myEvent = &shared_dto.Event{}
			err = json.Unmarshal(ev.Value, myEvent)
			if err != nil {
				fmt.Errorf("%s\n", err.Error())
			}

			shared_utils.PrettyDisplay("event", myEvent)

			selector := selector.NewSelector(myEvent.Scenario, *myEvent)
			res, err := selector.Scenario.Run()

			if err != nil {
				log.Fatal("-> %s \n", err.Error())
			}

			fmt.Println("RES : ", res)
		}
		fmt.Println("listen ")
	}

	kh.c.Close()
}
