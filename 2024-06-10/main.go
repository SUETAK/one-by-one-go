package main

import (
	"log"
	"one-by-one-go/2024-06-10/server"
)

func main() {
	// 引数が可変じゃないので、複数の値を入れられない
	svr := server.New("localhost", 8888)

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
