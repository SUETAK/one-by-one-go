package service

import "2023-10-07/domain"

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

func (c *countdownService) Start() {}
