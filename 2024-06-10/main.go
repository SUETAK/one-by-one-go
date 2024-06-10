package main

import (
	"log"
	"one-by-one-go/2024-06-10/server"
	"os"
	"time"
)

func main() {
	l, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	logger := log.New(l, "", log.LstdFlags)

	// 可変長引数にserver から呼び出したlogger, timeout を作成する関数を呼び出す
	// WithLogger, WithTimeout は引数をserver のフィールドに対して上書きする関数を返す。
	svr := server.New("localhost", 8888, server.WithLogger(logger), server.WithTimeout(time.Minute))

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
