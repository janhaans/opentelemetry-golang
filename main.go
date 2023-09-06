package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/janhaans/opentelemetry-golang/app"
)

func main() {
	ctx := context.Background()
	i := os.Stdin
	l := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	a := app.NewApp(ctx, i, l)
	err := a.Run(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
