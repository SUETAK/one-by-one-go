package domain

import (
	"context"
	"github.com/eiannone/keyboard"
)

type KeyMonitor struct {
	ch chan rune
}

func NewKeyMonitor() *KeyMonitor {
	// キーボードの監視を開始
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	return &KeyMonitor{
		ch: make(chan rune),
	}
}

func (k *KeyMonitor) Open(ctx context.Context) {
	// context からの終了通知を受け取って終了できるようにする
	go func() {
		for {
			select {
			case <-ctx.Done():
				err := keyboard.Close()
				if err != nil {
					panic(err)
				}
				return
			default:
				char, _, err := keyboard.GetSingleKey()
				if err != nil {
					panic(err)
				}
				k.ch <- char
			}
		}
	}()
}

func (k *KeyMonitor) GetCh() <-chan rune {
	return k.ch
}
