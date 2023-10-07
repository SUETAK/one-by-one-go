package main

import (
	"2023-10-07/service"
)

func main() {
	countdownService := service.NewCountdownService(10)
	countdownService.Start()
}
