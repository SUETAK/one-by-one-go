package service

import (
	"2023-10-07/domain"
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
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

	fmt.Printf("%d秒間のカウントダウンを設定しました。\n", c.timer.GetDuration())

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

	timerRunning := false

	for c.timer.GetDuration() >= 0 {
		select {
		case char := <-keyPressCh:
			switch char {
			case 't':
				timerRunning = true
			case 's':
				timerRunning = false
				fmt.Println("カウントダウンを停止しました")
			case 'r':
				timerRunning = true
				fmt.Println("カウントダウンを再開しました")
			default:
				fmt.Printf("'%c' キーが押されました。\n", char)
			}
		default:
			if timerRunning {
				c.timer.Countdown()
				fmt.Printf("%d秒\n", c.timer.GetDuration())
				time.Sleep(1 * time.Second)
			}
		}
	}

	fmt.Println("カウントダウン終了!")
}
