package main

import (
	"2023-10-07/service"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// サービス全体で使用するコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	// main 関数が終了するとコンテキストは終了する
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
	// コンテキストはメソッドに渡すのが良い
	countdownService.Start(ctx)
}
