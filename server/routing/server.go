package routing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shared/config"
	"syscall"
	"time"
)

func Serve() {
	configs := config.Get()
	r := GetRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", configs.Server.Port),
		Handler: r,
	}

	go func() {
		// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("listen: %s\n", err)
		// }
		fmt.Printf("====> listening to port : %s\n", configs.Server.Port)
		http.ListenAndServe(fmt.Sprintf(":%s", configs.Server.Port), mux)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")
}
