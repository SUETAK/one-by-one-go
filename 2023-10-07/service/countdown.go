package service

import (
	"2023-10-07/domain"
	"fmt"
	"github.com/eiannone/keyboard"
)

type CountdownService interface {
	Start()
}

func NewCountdownService(count int64) CountdownService {
	return &countdownService{
		timer: domain.NewTimer(count),
	}
}

type countdownService struct {
	timer *domain.Timer
}

func (c *countdownService) Start() {
	// キーボードの監視を開始
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Printf("%d秒間のカウントダウンを設定しました。カウントダウンを開始するには't'キーを押してください。\n", c.timer.GetDuration())

	// goroutineでキーイベントの監視を開始
	keyPressCh := make(chan rune)
	go func() {
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			keyPressCh <- char
		}
	}()

	for c.timer.GetDuration() >= 0 {
		char := <-keyPressCh
		if char != 0 {
			fmt.Printf("'%c' キーが押されました。\n", char)
			if char == 't' {
				c.timer.CountStart()
			}
			if char == 's' {
				c.timer.CountStop()
				fmt.Printf("カウントダウンを停止しました")
			}
			if char == 'r' {
				fmt.Printf("カウントダウンを再開しました")
				c.timer.CountStart()
			}
		}

		if c.timer.GetState() == "start" {
			c.timer.Countdown()
			fmt.Printf("%d秒\n", c.timer.GetDuration())
		}
	}

	fmt.Println("カウントダウン終了!")
}
