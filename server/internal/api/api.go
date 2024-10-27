package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shared/config"
	"syscall"
)

func RunApi(mux *http.ServeMux) {
	config := config.Get()

	go func() {
		fmt.Printf("====> listening to port : %d\n", config.Port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
		if err != nil {
			log.Fatal("error trying to launch the server", err.Error())
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

}
