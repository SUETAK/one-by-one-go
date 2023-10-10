package service

import (
	"2023-10-07/domain"
	"context"
	"fmt"
)

type CountdownService interface {
	Start(ctx context.Context)
}

func NewCountdownService(count int64) CountdownService {
	return &countdownService{
		timer: domain.NewTimer(count),
	}
}

type countdownService struct {
	timer *domain.Timer
}

func (c *countdownService) Start(ctx context.Context) {

	fmt.Printf("%d秒間のカウントダウンを設定しました。\n カウントを開始するためには't'キーを押してください。 カウントを中止する時は's'、再開する時は'r'キーを押してください", c.timer.GetDuration())

	keyMonitor := c.timer.GetKeyMonitor()
	keyMonitor.Open(ctx)

	for c.timer.GetDuration() > 0 {
		select {
		case char := <-keyMonitor.GetCh():
			switch char {
			case 't':
				c.timer.CountStart()
				fmt.Println("カウントダウンを開始しました")
			case 's':
				c.timer.CountStop()
				fmt.Println("カウントダウンを停止しました")
			case 'r':
				c.timer.CountStart()
				fmt.Println("カウントダウンを再開しました")
			default:
				fmt.Printf("'%c' キーが押されました。\n", char)
			}
		case <-ctx.Done():
			fmt.Println("カウントダウンが中断されました")
		default:
			if c.timer.GetState() == "start" {
				c.timer.Countdown()
				fmt.Printf("%d秒\n", c.timer.GetDuration())
			}
		}
	}

	fmt.Println("カウントダウン終了!")
}
