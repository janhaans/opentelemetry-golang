package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/janhaans/opentelemetry-golang/app"
)

func main() {
	ctx := context.Background()
	i := os.Stdin
	l := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	a := app.NewApp(ctx, i, l)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)

	go func() {
		errCh <- a.Run(ctx)
	}()

	select {
	case <-sigCh:
		l.Println("\nGoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatalf("%v\n", err)
		}
	}
}
