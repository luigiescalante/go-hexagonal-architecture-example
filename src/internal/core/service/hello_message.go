package service

import (
	"go.hex-architecture/internal/core/domain"
	"go.hex-architecture/internal/core/ports"
)

type HelloMessageSrv struct {
	Repository ports.HelloMessageRepo
	HelloMsg   domain.HelloMessage
}

func NewHelloMessage(repo ports.HelloMessageRepo) *HelloMessageSrv {
	return &HelloMessageSrv{
		Repository: repo,
	}
}

func (s *HelloMessageSrv) CreateMessage(name string) {
	s.HelloMsg.Message = "Hello " + name + " !!!"
}

func (s *HelloMessageSrv) GetMessage() string {
	return s.HelloMsg.Message
}
