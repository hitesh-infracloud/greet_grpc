package main

import (
	"github.com/hitesh-infracloud/awesome-code/greet_app/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a := app.GetNewGreetApp()
	a.Start()
	<-interrupt()
	a.Shutdown()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
