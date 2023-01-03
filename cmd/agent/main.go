package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	m := newMxMetrics()
	client := newHttpClient()

	var wg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		updateMetrics(ctx, m)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendMetrics(ctx, m, client)
	}()

	wg.Wait()
}
