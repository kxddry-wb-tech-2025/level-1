package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NWorkers запускает N воркеров, которые просто ждут закрытия контекста
func NWorkers(ctx context.Context, n int) {
	for range n {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
			}
		}()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	NWorkers(ctx, 3)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)
	<-stop // при получении сигнала Ctrl+C отменяем контекст, все воркеры дохнут
	cancel()
}
