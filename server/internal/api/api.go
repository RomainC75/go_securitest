package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func RunApi(mux *http.ServeMux) {
	port := viper.Get("SERVER_PORT")

	go func() {
		fmt.Printf("====> listening to port : %s\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
		if err != nil {
			log.Fatal("error trying to launch the server", err.Error())
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

}
