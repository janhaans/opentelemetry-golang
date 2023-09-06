package app

import (
	"context"
	"fmt"
	"io"
	"log"
)

type App struct {
	i io.Reader
	l *log.Logger
}

func NewApp(ctx context.Context, i io.Reader, l *log.Logger) *App {
	return &App{i, l}
}

func (a *App) Run(ctx context.Context) error {
	input, err := a.Poll(ctx)
	if err != nil {
		return err
	}
	result, err := fibonacci(input)
	if err != nil {
		return err
	}
	a.Write(ctx, input, result)
	return nil
}

func (a *App) Poll(ctx context.Context) (int, error) {
	fmt.Print("What Fibonacci number would you like to know: ")

	var input int
	_, err := fmt.Fscanf(a.i, "%d", &input)
	return input, err
}

func (a *App) Write(ctx context.Context, input, result int) {
	a.l.Printf("The Fibonacci number of %d = %d", input, result)
}

func fibonacci(n int) (int, error) {
	if n < 0 {
		err := fmt.Errorf("fibonacci number is smaller then 0")
		return 0, err
	}
	if n == 0 || n == 1 {
		return n, nil
	}
	n1 := 0
	n2 := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = n1 + n2
		n1 = n2
		n2 = result
	}
	return result, nil

}
