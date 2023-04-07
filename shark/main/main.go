package main

import (
	"context"
	"mp-projeto/shared/resource"
	"mp-projeto/shark/adapter"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		cancel()
	}()

	sr := resource.NewServerResource("env.toml")

	//Register Handlers
	sr.AddHandler(adapter.NewIncidenceHandlerRest(sr))

	sr.Run(ctx)
}
