package service

type CountdownService interface {
	Start()
}

func NewCountdownService() CountdownService {
	return &countdownService{}
}

type countdownService struct {
	count int64
}

func (c *countdownService) Start() {}
