package domain

import "github.com/eiannone/keyboard"

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

func (k *KeyMonitor) Open() {
	go func() {
		for {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			k.ch <- char
		}
	}()
}

func (k *KeyMonitor) Stop() {
	defer keyboard.Close()
}

func (k *KeyMonitor) GetCh() chan rune {
	return k.ch
}
