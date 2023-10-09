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

	fmt.Println("10秒間のカウントダウンを開始します。途中で任意のキーを押してください。")

	// タイマーを設定 (例: 10秒)
	timerDuration := c.timer.GetDuration()
	endTime := time.Now().Add(timerDuration)

	for time.Now().Before(endTime) {
		// キーイベントの監視 (非ブロッキング)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if char != 0 {
			fmt.Printf("'%c' キーが押されました。\n", char)
		}

		// 一定間隔（例: 0.1秒）で確認
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("カウントダウン終了!")
}
