package main

import (
	"2023-10-07/service"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// OSのシグナルを捕捉するためのチャネルの作成
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// シグナルが受信された場合にcontextをキャンセル
	go func() {
		<-sigCh
		cancel()
	}()

	countdownService := service.NewCountdownService(10)
	countdownService.Start(ctx)
}
