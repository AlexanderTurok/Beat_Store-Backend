package service

type Authorization interface {
}

type Beat interface {
}

type Service struct {
	Authorization
	Beat
}

func NewService() *Service {
	return &Service{}
}
